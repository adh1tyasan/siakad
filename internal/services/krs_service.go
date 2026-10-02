package services

import (
	"errors"
	"siakad/internal/models"
	"siakad/internal/repositories"
)

type TambahKRSRequest struct {
	MahasiswaID     uint `json:"mahasiswa_id" binding:"required"`
	TahunAkademikID uint `json:"tahun_akademik_id" binding:"required"`
	JadwalID        uint `json:"jadwal_id" binding:"required"`
}

func TambahJadwalKeKRS(req TambahKRSRequest) error {
	// 1. Cek apakah Jadwal yang dipilih ada
	jadwalBaru, err := repositories.GetJadwalBerdasarkanID(req.JadwalID)
	if err != nil {
		return errors.New("Jadwal kuliah tidak ditemukan")
	}

	// 2. Cari KRS Draft milik mahasiswa. Jika belum ada, buatkan baru.
	krs, err := repositories.GetDraftKRSMahasiswa(req.MahasiswaID, req.TahunAkademikID)
	if err != nil {
		// Buat KRS baru dengan status Draft
		krs = &models.KRS{
			MahasiswaID:     req.MahasiswaID,
			TahunAkademikID: req.TahunAkademikID,
			Status:          "Draft",
			TotalSKS:        0,
		}
		if errCreate := repositories.CreateKRSBaru(krs); errCreate != nil {
			return errors.New("Gagal membuat data KRS baru" + errCreate.Error())
		}
	}

	// 3. Validasi Maksimal SKS (Batas 24 SKS)
	sksMataKuliahBaru := jadwalBaru.MataKuliah.SKS
	if krs.TotalSKS+sksMataKuliahBaru > 24 {
		return errors.New("SKS melebihi batas maksimal (24 SKS)")
	}

	// 4. Validasi Bentrok Jadwal & Duplikasi
	for _, detail := range krs.KRSDetail {
		// Cek Duplikasi: Apakah mata kuliah ini sudah diambil?
		if detail.JadwalID == req.JadwalID {
			return errors.New("Mata kuliah ini sudah ada di dalam KRS Anda")
		}

		// Cek Bentrok: Hari sama dan Jam Mulai sama (Logika sederhana)
		jadwalLama := detail.Jadwal
		if jadwalLama.Hari == jadwalBaru.Hari && jadwalLama.JamMulai == jadwalBaru.JamMulai {
			return errors.New("Jadwal bentrok dengan mata kuliah lain di KRS Anda")
		}
	}

	// 5. Jika lolos semua validasi, simpan ke KRS Detail
	krsDetail := models.KRSDetail{
		KRSID:    krs.ID,
		JadwalID: req.JadwalID,
	}

	totalSKSBaru := krs.TotalSKS + sksMataKuliahBaru
	return repositories.TambahDetailKRS(krs.ID, &krsDetail, totalSKSBaru)
}

// Mahasiswa Mengunci draft dan mengajukan Ke Dosen PA
func AjukanKRS(krsID uint) error {
	krs, err := repositories.GetKRSByID(krsID)
	if err != nil {
		return errors.New("KRS Tidak ditemukan")
	}
	if krs.Status != "Draft" {
		return errors.New("Hanya KRS yang berstatus draft yang dapat diajuka")
	}
	return repositories.UpdateStatusKRS(krsID, "Diajukan", "")
}

// Persetujuan Tahap1  : Dosen PA
func PersetujuanPA(krsID uint, status, catatan string) error {
	krs, err := repositories.GetKRSByID(krsID)
	if err != nil {
		return errors.New("KRS tidak ditemukan")
	}
	if krs.Status != "Diajukan" {
		return errors.New("KRS Belum diajukan Mahasiswa atau Sudah diproses")
	}
	if krs.Status != "Disetujui PA" && krs.Status != "Ditolah" {
		return errors.New("Status tidak valid. Gunakan 'Disetujui PA' atau 'Ditolak'")
	}
	return repositories.UpdateStatusKRS(krsID, status, catatan)
}

// Persetujuan Tahap 2 : Ketua Prodi
func PersetujuanKaprodi(krsID uint, status, catatan string) error {
	krs, err := repositories.GetKRSByID(krsID)
	if err != nil {
		return errors.New("KRS tidak ditemukan")
	}
	if krs.Status != "Disetujui PA" {
		return errors.New("KRS Harus disetujui Dosen PA terlebih dahulu")
	}
	if krs.Status != "Disetujui Kaprodi" && krs.Status != "Ditolal" {
		return errors.New("Status tidak valid. Gunakan 'Disetujui Kaprodi' atau 'Ditolak'")
	}
	return repositories.UpdateStatusKRS(krsID, status, catatan)
}
