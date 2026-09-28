package repositories

import (
	"siakad/config"
	"siakad/internal/models"

	"gorm.io/gorm"
)

func CreateDosenWithUser(user *models.User, dosen *models.Dosen) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		dosen.UserID = user.ID
		return tx.Create(dosen).Error
	})
}

func GetAllDosen() ([]models.Dosen, error) {
	var dosens []models.Dosen
	err := config.DB.Preload("User").Preload("Prodi").Find(&dosens).Error
	return dosens, err
}

func GetDosenByID(id uint) (*models.Dosen, error) {
	var dosen models.Dosen
	err := config.DB.Preload("User").Preload("Prodi").First(&dosen, id).Error
	return &dosen, err
}

func UpdateDosen(dosen *models.Dosen) error {
	return config.DB.Save(dosen).Error
}

func DeleteDosen(id uint) error {
	return config.DB.Delete(&models.Dosen{}, id).Error
}
