package repositories

import (
	"siakad/config"
	"siakad/internal/models"
)

// Simpan data sesi pertemuan kelas
func CreatePertemuan(pertemuan *models.Pertemuan) error {
	return config.DB.Create(pertemuan).Error
}

// simpan banyak data kehadiran sekaligus (Bulk Insert)
func CreateAbsensi(absensi []models.Absensi) error {
	return config.DB.Create(&absensi).Error
}

// Mengambil data pertemuan untuk dicek apakah ada
func GetPertemuanByID(id uint) (*models.Pertemuan, error) {
	var pertemuan models.Pertemuan
	err := config.DB.First(&pertemuan, id).Error
	return &pertemuan, err
}

// Simpan perubahan data pertemuan seperti tanggal atau materi
func UpdatePertemuan(pertemuan *models.Pertemuan) error {
	return config.DB.Save(pertemuan).Error
}

// Update status absensi mahasiswa di pertemuan tertentu
func UpdateStatusAbsensi(pertemuanID uint, mahasiswaID uint, status string) error {
	//Update Spesifik ke mahasiswa tertentu di pertemuan
	return config.DB.Model(&models.Absensi{}).Where("pertemuan_id = ? AND mahasiswa_id = ?", pertemuanID, mahasiswaID).Update("status", status).Error
}

func GetAbsensiByMahasiswaID(mahasiswaID uint) ([]models.Absensi, error) {
	var listAbsensi []models.Absensi
	err := config.DB.
		Preload("Pertemuan").
		Preload("Pertemuan.Jadwal").
		Preload("Pertemuan.Jadwal.MataKuliah").
		Where("mahasiswa_id = ? ", mahasiswaID).
		Find(&listAbsensi).Error
	return listAbsensi, err
}
