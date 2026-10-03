package handlers

import (
	"net/http"
	"siakad/internal/repositories"
	"siakad/internal/services"
	"siakad/internal/utils"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest

	//Validasi Input JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "format request tidak valid",
		})
		return
	}

	//Cari user berdasarkan username
	user, err := repositories.GetUserByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"succes":  false,
			"message": "username atau password salah",
		})
		return
	}

	//verifikasi password
	if !utils.CheckPassswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"succes":  false,
			"message": "username atau password salah",
		})
		return
	}

	//Create Token JWT
	token, err := utils.GenerateToken(user.ID, user.Role.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"succes":  false,
			"message": "Failed Create Token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "Login Berhasil",
		"data": gin.H{
			"token": token,
			"role":  user.Role.Name,
		},
	})
}

func Register(ctx *gin.Context) {
	var req services.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "Format salah. Pastikan menyertakan username, password (min 6 karakter), dan role_id",
		})
		return
	}
	if err := services.RegisterUser(req); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Registrasi berhasil, silakan login",
	})
}
