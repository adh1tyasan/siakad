package repositories

import (
	"siakad/config"
	"siakad/internal/models"
)

func CreateMataKuliah(mk *models.MataKuliah) error {
	return config.DB.Create(mk).Error
}

func GetAllMataKuliah() ([]models.MataKuliah, error) {
	var mks []models.MataKuliah
	//Preload Prodi untuk mengetahui mata kuliah milik prodi mana
	err := config.DB.Preload("Prodi").Find(&mks).Error
	return mks, err
}

func GetMataKuliahByID(id uint) (*models.MataKuliah, error) {
	var mk models.MataKuliah
	err := config.DB.Preload("Prodi").First(&mk, id).Error
	return &mk, err
}

func UpdateMataKuliah(mk *models.MataKuliah) error {
	return config.DB.Save(mk).Error
}

func DeleteMataKuliah(id uint) error {
	return config.DB.Delete(&models.MataKuliah{}, id).Error
}
