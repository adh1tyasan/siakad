package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Generate Token membuat JWT untuk user yang berhasil login
func GenerateToken(userID uint, role string) (string, error) {
	//Token berlaku selama 24 jam
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Validate token untuk cek keaslian token
func ValidateToken(tokenString string) (*jwt.Token, error) {
	//Jika JWT secret belum diload(karena init file), ambil lagi
	if len(jwtSecret) == 0 {
		jwtSecret = []byte(os.Getenv("JWT_SECREt"))
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("metode signing tidak valid")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, err
}
