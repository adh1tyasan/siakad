package services

import (
	"errors"
	"siakad/internal/models"
	"siakad/internal/repositories"
)

// struct untuk menangkap 1 data mahasiswa (di dalam aarray)
type AbsensiDetailReq struct {
	MahasiswaID uint   `json:"mahasiswa_id" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

// Struct utama yang akan ditangkap dari POstman
type InputAbsensiRequest struct {
	JadwalID    uint               `json:"jadwal_id" binding:"required"`
	PertemuanKe int                `json:"pertemuan_ke" binding:"required"`
	Tanggal     string             `json:"tanggal" binding:"required"`
	Materi      string             `json:"materi" binding:"required"`
	DataAbsem   []AbsensiDetailReq `json:"data_absen" binding:"required"`
}

type RekapAbsensiResponse struct {
	MataKuliah     string `json:"mata_kuliah"`
	Hadir          int    `json:"hadir"`
	Izin           int    `json:"izin"`
	Sakit          int    `json:"sakit"`
	Alpa           int    `json:"alpa"`
	TotalPertemuan int    `json:"total_pertemuan"`
}

type UpdateAbsensiRequest struct {
	Tanggal   string             `json:"tanggal" binding:"required"`
	Materi    string             `json:"materi" binding:"required"`
	DataAbsen []AbsensiDetailReq `json:"data_absen" binding:"required"`
}

func ProsesInputAbsensi(req InputAbsensiRequest) error {
	//Buat data pertemuan terlebih dahulu
	pertemuan := models.Pertemuan{
		JadwalID:    req.JadwalID,
		PertemuanKe: req.PertemuanKe,
		Tanggal:     req.Tanggal,
		Materi:      req.Materi,
	}

	if err := repositories.CreatePertemuan(&pertemuan); err != nil {
		return errors.New("Gagal Menyimpan Data Pertemuan: " + err.Error())
	}

	//Siapkan list data absensi menggunakn ID pertemuan yang baru dibuat
	var listAbsensi []models.Absensi
	for _, item := range req.DataAbsem {
		listAbsensi = append(listAbsensi, models.Absensi{
			PertemuanID: pertemuan.ID,
			MahasiswaID: item.MahasiswaID,
			Status:      item.Status,
		})
	}

	//Simpan ke tabel absensi sekaligus
	if err := repositories.CreateAbsensi(listAbsensi); err != nil {
		return errors.New("Gagal Menyimpan Detail Absensi: " + err.Error())
	}
	return nil
}

func ProsesUpdateAbsensi(pertemuanID uint, req UpdateAbsensiRequest) error {
	//Pertama di cek apakah pertemuan valid
	pertemuan, err := repositories.GetPertemuanByID(pertemuanID)
	if err != nil {
		return errors.New("Data pertemuan tidak ditemukan")
	}

	//lanjut update materi dan tanggal
	pertemuan.Tanggal = req.Tanggal
	pertemuan.Materi = req.Materi
	if err := repositories.UpdatePertemuan(pertemuan); err != nil {
		return errors.New("Gagal memperbarui data pertemuan: " + err.Error())
	}

	//update status absen mahasiswa
	for _, item := range req.DataAbsen {
		errUpdate := repositories.UpdateStatusAbsensi(pertemuanID, item.MahasiswaID, item.Status)
		if errUpdate != nil {
			return errors.New("Gagal memperbarui abasensi mahasiswa")
		}
	}
	return nil
}

func GetRekapAbsensiMahasiswa(userID uint) ([]RekapAbsensiResponse, error) {
	//cari profile mahasiswa Berdasarkan token login
	mhs, err := repositories.GetMahasiswaByUserID(userID)
	if err != nil {
		return nil, errors.New("Data mahasiswa tidak ditemukan")
	}

	//Tarik semua data absensi nya
	listAbsensi, err := repositories.GetAbsensiByMahasiswaID(mhs.ID)
	if err != nil {
		return nil, errors.New("Gagal mengambil data absensi: " + err.Error())
	}

	// Kelompokkan dan hitung per matakuliah pakai map
	rekapMap := make(map[uint]*RekapAbsensiResponse)

	for _, absen := range listAbsensi {
		mk := absen.Pertemuan.Jadwal.MataKuliah
		mkID := mk.ID

		//jika mata kuliah tidak ada di map, maka buatkan kerangkanya
		if _, exist := rekapMap[mkID]; !exist {
			rekapMap[mkID] = &RekapAbsensiResponse{
				MataKuliah: mk.Nama,
			}
		}
		rekapMap[mkID].TotalPertemuan++
		switch absen.Status {
		case "Hadir":
			rekapMap[mkID].Hadir++
		case "Izin":
			rekapMap[mkID].Izin++
		case "Sakit":
			rekapMap[mkID].Sakit++
		case "Alpa":
			rekapMap[mkID].Alpa++
		}
	}

	//Ubah format map menjadi Array/Slice agar bagus saat jadi JSON
	var result []RekapAbsensiResponse
	for _, v := range rekapMap {
		result = append(result, *v)
	}
	return result, nil
}
