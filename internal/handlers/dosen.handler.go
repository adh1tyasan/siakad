package handlers

import (
	"net/http"
	"siakad/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DosenInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	NIDN     string `json:"nidn" binding:"required"`
	Nama     string `json:"nama" binding:"required"`
	ProdiID  uint   `json:"prodi_id" binding:"required"`
}

type DosenUpdateInput struct {
	NIDN    string `json:"nidn" binding:"required"`
	Nama    string `json:"nama" binding:"required"`
	ProdiID uint   `json:"prodi_id" binding:"required"`
}

func CreateDosen(ctx *gin.Context) {
	var input DosenInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "format input salah",
		})
		return
	}

	if err := services.CreateDosen(input.Username, input.Password, input.NIDN, input.Nama, input.ProdiID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"succes":  true,
		"message": "Dosen dan Akun User berhasil dibuat",
	})
}

func GetAllDosen(ctx *gin.Context) {
	dosens, err := services.GetAllDosen()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"succes":  false,
			"message": "Gagal mengambil data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes": true,
		"data":   dosens,
	})
}

func UpdateDosen(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var input DosenUpdateInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "format input salah",
		})
		return
	}

	if err := services.UpdateDosen(uint(id), input.NIDN, input.Nama, input.ProdiID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "Data Dosen Berhasil diperbarui",
	})
}

func DeleteDosen(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := services.DeleteDosen(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "Data dosen berhasil dihapus",
	})
}
