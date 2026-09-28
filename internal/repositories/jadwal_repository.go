package repositories

import (
	"siakad/config"
	"siakad/internal/models"
)

// Pengecekan data Master
func IsRuanganExists(id uint) bool {
	var count int64
	config.DB.Model(&models.Ruangan{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func IsMataKuliahExists(id uint) bool {
	var count int64
	config.DB.Model(&models.MataKuliah{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func IsDosenExists(id uint) bool {
	var count int64
	config.DB.Model(&models.Dosen{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func IsKelasExists(id uint) bool {
	var count int64
	config.DB.Model(&models.Kelas{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func IsTahunAkademikExists(id uint) bool {
	var count int64
	config.DB.Model(&models.TahunAkademik{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func CreateJadwal(jadwal *models.Jadwal) error {
	return config.DB.Create(jadwal).Error
}

func GetAllJadwal() ([]models.Jadwal, error) {
	var jadwals []models.Jadwal
	//Preload semua relasi agar response JSON lengkap
	err := config.DB.Preload("TahunAkademik").Preload("MataKuliah").Preload("Kelas").Preload("Dosen").Preload("Ruangan").Find(&jadwals).Error
	return jadwals, err
}

func CekJadwalBentrok(ruanganID uint, hari, jamMulai string) bool {
	var count int64
	//cek apakah di hari, ruangan dan jam mulai yang sama sudah ada jadwal
	config.DB.Model(&models.Jadwal{}).Where("ruangan_id = ? AND hari = ? AND jam_mulai = ?", ruanganID, hari, jamMulai).Count(&count)
	return count > 0
}

func GetJadwalByID(id uint) (*models.Jadwal, error) {
	var jadwal models.Jadwal
	err := config.DB.Preload("TahunAkademik").Preload("MataKuliah").Preload("Kelas").Preload("Dosen").Preload("Ruangan").First(&jadwal, id).Error
	return &jadwal, err
}

func UpdateJadwal(jadwal *models.Jadwal) error {
	return config.DB.Save(jadwal).Error
}

func DeleteJadwal(id uint) error {
	return config.DB.Delete(&models.Jadwal{}, id).Error
}

// CekJadwalBentrokExcludeID untuk mengecek bentrok saat Update, abaikan ID jadwal yang sedang di-edit
func CekJadwalBentrokExcludeID(ruanganID uint, hari, jamMulai string, excludeID uint) bool {
	var count int64
	config.DB.Model(&models.Jadwal{}).Where("ruangan_id = ? AND hari = ? AND jam_mulai = ? AND id != ?", ruanganID, jamMulai, excludeID).Count(&count)
	return count > 0
}
