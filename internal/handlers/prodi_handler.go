package handlers

import (
	"net/http"
	"siakad/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProdiInput struct {
	Kode string `json:"kode" binding:"required"`
	Nama string `json:"nama" binding:"required"`
}

func CreateProdi(ctx *gin.Context) {
	var input ProdiInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "format input salah",
		})
		return
	}
	if err := services.CreateProdi(input.Kode, input.Nama); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"succes":  true,
		"message": "Prodi berhasil ditambahkan",
	})
}

func GetAllProdi(ctx *gin.Context) {
	prodis, err := services.GetAllProdi()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"succes":  false,
			"message": "Gagal mengambil data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes": true,
		"data":   prodis,
	})
}

func UpdateProdi(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var input ProdiInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "Format input salah",
		})
		return
	}
	if err := services.UpdateProdi(uint(id), input.Kode, input.Nama); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "Prodi Berhasil diupdate",
	})
}

func DeleteProdi(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := services.DeleteProdi(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "Prodi Berhasil dihapus",
	})
}
