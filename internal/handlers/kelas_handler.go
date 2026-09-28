package handlers

import (
	"net/http"
	"siakad/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KelasInput struct {
	Nama            string `json:"nama" binding:"required"`
	ProdiID         uint   `json:"prodi_id" binding:"required"`
	TahunAkademikID uint   `json:"tahun_akademik_id" binding:"required"`
}

func CreateKelas(ctx *gin.Context) {
	var input KelasInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "Format input salah",
		})
		return
	}

	if err := services.CreateKelas(input.Nama, input.ProdiID, input.TahunAkademikID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"succes":  true,
		"message": "Kelas berhasil ditambahkan",
	})
}

func GetAllKelas(ctx *gin.Context) {
	kelas, err := services.GetAllKelas()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"succes":  false,
			"message": "Gagal Mengambil data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"succes": true,
		"data":   kelas,
	})
}

func UpdateKelas(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var input KelasInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "format input salah",
		})
		return
	}

	if err := services.UpdateKelas(uint(id), input.Nama, input.ProdiID, input.TahunAkademikID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "Kelas Berhasil diperbarui",
	})
}

func DeleteKelas(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := services.DeleteKelas(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "Kelas berhasil di hapus",
	})
}
