package handlers

import (
	"net/http"
	"siakad/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MahasiswaInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	NPM      string `json:"nim" binding:"required"`
	Nama     string `json:"nama" binding:"required"`
	ProdiID  uint   `json:"prodi_id" binding:"required"`
	Kelas    string `json:"kelas" binding:"required"`
}

type MahasiswaUpdateInput struct {
	NPM     string `json:"nim" binding:"required"`
	Nama    string `json:"nama" binding:"required"`
	ProdiID uint   `json:"prodi_id" binding:"required"`
	Kelas   string `json:"kelas" binding:"required"`
}

func CreateMahasiswa(ctx *gin.Context) {
	var input MahasiswaInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "Format input salah atau ada Data Kosong",
		})
		return
	}

	err := services.CreateMahasiswa(input.Username, input.Password, input.NPM, input.Nama, input.Kelas, input.ProdiID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"succes":  true,
		"message": "Mahasiswa dan Akun User berhasil dibuat",
	})
}

func GetAllMahasiswa(ctx *gin.Context) {
	mahasiswas, err := services.GetAllMahasiswa()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"succes":  false,
			"message": "Gagal mengambil data Mahasiswa",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes": true,
		"data":   mahasiswas,
	})
}

func UpdateMahasiswa(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var input MahasiswaUpdateInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "format input salah",
		})
		return
	}

	if err := services.UpdateMahasiswa(uint(id), input.NPM, input.Nama, input.Kelas, input.ProdiID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "Data Mahasiswa Berhasil diperbarui",
	})
}

func DeleteMahasiswa(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := services.DeleteMahasiswa(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "Data Mahasiswa Berhasil dihapus",
	})
}
