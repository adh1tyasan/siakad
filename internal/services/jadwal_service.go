package services

import (
	"errors"
	"siakad/internal/models"
	"siakad/internal/repositories"
)

func CreateJadwal(taID, mkID, kelasID, dosenID, ruanganID uint, hari, jamMulai, jamSelesai string) error {
	// 1. Validasi Keberadaan Data Master
	if !repositories.IsTahunAkademikExists(taID) {
		return errors.New("Tahun Akademik tidak ditemukan")
	}
	if !repositories.IsMataKuliahExists(mkID) {
		return errors.New("Mata Kuliah tidak ditemukan")
	}
	if !repositories.IsKelasExists(kelasID) {
		return errors.New("Kelas tidak ditemukan")
	}
	if !repositories.IsDosenExists(dosenID) {
		return errors.New("Dosen tidak ditemukan")
	}
	if !repositories.IsRuanganExists(ruanganID) {
		return errors.New("Ruangan tidak ditemukan")
	}

	// 2. Validasi Bentrok Jadwal
	if repositories.CekJadwalBentrok(ruanganID, hari, jamMulai) {
		return errors.New("Jadwal Bentrok: ruangan sudah dipakai pada hari dan jam tersebut")
	}

	jadwal := &models.Jadwal{
		TahunAkademikID: taID,
		MataKuliahID:    mkID,
		KelasID:         kelasID,
		DosenID:         dosenID,
		RuanganID:       ruanganID,
		Hari:            hari,
		JamMulai:        jamMulai,
		JamSelesai:      jamSelesai,
	}
	return repositories.CreateJadwal(jadwal)
}

func GetAllJadwal() ([]models.Jadwal, error) {
	return repositories.GetAllJadwal()
}

func UpdateJadwal(id, taID, mkID, kelasID, dosenID, ruanganID uint, hari, jamMulai, jamSelesai string) error {
	jadwal, err := repositories.GetJadwalByID(id)
	if err != nil {
		return errors.New("data jadwal tidak ditemukan")
	}

	// 1. Validasi Keberadaan Data Master
	if !repositories.IsTahunAkademikExists(taID) {
		return errors.New("Tahun Akademik tidak ditemukan")
	}
	if !repositories.IsMataKuliahExists(mkID) {
		return errors.New("Mata Kuliah tidak ditemukan")
	}
	if !repositories.IsKelasExists(kelasID) {
		return errors.New("Kelas tidak ditemukan")
	}
	if !repositories.IsDosenExists(dosenID) {
		return errors.New("Dosen tidak ditemukan")
	}
	if !repositories.IsRuanganExists(ruanganID) {
		return errors.New("Ruangan tidak ditemukan")
	}

	// 2. Validasi bentrok (abaikan id Jadwal sendiri)
	if repositories.CekJadwalBentrokExcludeID(ruanganID, hari, jamMulai, id) {
		return errors.New("Jadwal Bentrok: ruangan sudah dipakai pada hari dan jam tersebut")
	}

	jadwal.TahunAkademikID = taID
	jadwal.MataKuliahID = mkID
	jadwal.KelasID = kelasID // Ditambahkan karena sebelumnya terlewat
	jadwal.DosenID = dosenID
	jadwal.RuanganID = ruanganID
	jadwal.Hari = hari
	jadwal.JamMulai = jamMulai
	jadwal.JamSelesai = jamSelesai

	return repositories.UpdateJadwal(jadwal)
}

func DeleteJadwal(id uint) error {
	_, err := repositories.GetJadwalByID(id)
	if err != nil {
		return errors.New("data jadwal tidak ditemukan")
	}
	return repositories.DeleteJadwal(id)
}
