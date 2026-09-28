package handlers

import (
	"net/http"
	"siakad/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MataKuliahInput struct {
	Kode     string `json:"kode" binding:"required"`
	Nama     string `json:"nama" binding:"required"`
	SKS      int    `json:"sks" binding:"required,gt=0"` // SKS harus lebih dari 0
	Semester int    `json:"semester" binding:"required,gt=0"`
	ProdiID  uint   `json:"prodi_id" binding:"required"`
}

func CreataMataKuliah(ctx *gin.Context) {
	var input MataKuliahInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "format input salah",
		})
		return
	}

	err := services.CreateMataKuliah(input.Kode, input.Nama, input.SKS, input.Semester, input.ProdiID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"succes":  true,
		"message": "Mata Kuliah berhasil ditambahkan",
	})

}

func GetAllMataKuliah(ctx *gin.Context) {
	mks, err := services.GetAllMataKuliah()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"succes":  false,
			"message": "gagal mengambil data mata kuliah",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes": true,
		"data":   mks,
	})
}

func UpdateMataKuliah(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var input MataKuliahInput //Pakai MataKuliahINput yang sma karena fieldnya identik

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "Format input salah",
		})
		return
	}

	err := services.UpdateMataKuliah(uint(id), input.Kode, input.Nama, input.SKS, input.Semester, input.ProdiID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "Data Mata kuliah berhasil diperbarui",
	})
}

func DeleteMataKuliah(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := services.DeleteMataKuliah(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "Data Mata Kuliah berhasil dihapus",
	})
}
