package repositories

import (
	"siakad/config"
	"siakad/internal/models"
)

func CreateProdi(prodi *models.Prodi) error {
	return config.DB.Create(prodi).Error
}

func GetAllProdi() ([]models.Prodi, error) {
	var prodis []models.Prodi
	err := config.DB.Find(&prodis).Error
	return prodis, err
}

func GetProdiByID(id uint) (*models.Prodi, error) {
	var prodi models.Prodi
	err := config.DB.First(&prodi, id).Error
	return &prodi, err
}

func UpdateProdi(prodi *models.Prodi) error {
	return config.DB.Save(prodi).Error
}

func DeleteProdi(id uint) error {
	return config.DB.Delete(&models.Prodi{}, id).Error
}
