package repositories

import (
	"siakad/config"
	"siakad/internal/models"
)

func CreateKelas(kelas *models.Kelas) error {
	return config.DB.Create(kelas).Error
}

func GetAllKelas() ([]models.Kelas, error) {
	var kelas []models.Kelas
	//Pakai Preload agar data nama Prodi dan Tahun AKademik tampil
	err := config.DB.Preload("Prodi").Preload("TahunAkademik").Find(&kelas).Error
	return kelas, err
}

func GetKelasByID(id uint) (*models.Kelas, error) {
	var kelas models.Kelas
	err := config.DB.Preload("Prodi").Preload("TahunAkademik").First(&kelas, id).Error
	return &kelas, err
}

func UpdateKelas(kelas *models.Kelas) error {
	return config.DB.Save(kelas).Error
}

func DeleteKelas(id uint) error {
	return config.DB.Delete(&models.Kelas{}, id).Error
}
