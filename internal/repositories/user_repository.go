package repositories

import (
	"siakad/config"
	"siakad/internal/models"
)

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	// Preload("Role") berguna agar kita langsung tahu user ini admin, dosen, dll
	err := config.DB.Preload("Role").Where("username = ?", username).First(&user).Error
	return &user, err
}

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}
