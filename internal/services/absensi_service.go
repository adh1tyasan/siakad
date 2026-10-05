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
