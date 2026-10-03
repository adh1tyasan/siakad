package services

import (
	"errors"
	"siakad/internal/models"
	"siakad/internal/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LogiRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

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

func LoginUser(req LogiRequest) (string, error) {
	//Cari user berdasarkan username
	user, err := repositories.GetUserByUsername(req.Username)
	if err != nil {
		return "", errors.New("Username atau Password salah")
	}

	//cocokkan password yang diinput dengan password hash di database
	errHash := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if errHash != nil {
		return "", errors.New("Username atau Password salah")
	}

	//Create Token JWT if Succces
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role.Name,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Gunakan Secret Key rahasia (Idealnya ditaruh di file .env, tapi untuk sekarang kita hardcode)
	tokenString, errToken := token.SignedString([]byte("KODE_RAHASIA_SIAKAD_123"))
	if errToken != nil {
		return "", errors.New("Gagal membuat token keamanan")
	}

	return tokenString, nil
}
