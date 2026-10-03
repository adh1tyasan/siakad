package services

import (
	"errors"
	"siakad/internal/models"
	"siakad/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID   uint   `json:"role_id" binding:"required"`
}

func RegisterUser(req RegisterRequest) error {
	//Cek apakah username sudah ada di database
	_, err := repositories.GetUserByUsername(req.Username)
	if err == nil {
		//Jika err == nil maka usernamenya KETEMU, jadi tidak boleh dipakai lagi
		return errors.New("Username sudah terdaftar, Silahkan buat username lain")
	}
	//Hash Password
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if errHash != nil {
		return errors.New("Gagal mengenkripsi password")
	}

	//Data User baru
	userBaru := &models.User{
		Username: req.Username,
		Password: string(hashedPassword),
		RoleID:   req.RoleID,
	}

	//Request Ke repository untuk simpan ke database
	return repositories.CreateUser(userBaru)
}
