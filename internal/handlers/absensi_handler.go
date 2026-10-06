package handlers

import (
	"net/http"
	"siakad/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InputAbsensi(ctx *gin.Context) {
	var req services.InputAbsensiRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "Format data tidak valid, pastikan data absensi lengkap",
		})
		return
	}

	if err := services.ProsesInputAbsensi(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"succes":  true,
		"message": "Data pertemuan dan absensi berhasil disimpan",
	})
}

func UpdateAbsensi(ctx *gin.Context) {
	id, errParam := strconv.Atoi(ctx.Param("id"))
	if errParam != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "ID Pertemuan tidak valid",
		})
		return
	}

	var req services.UpdateAbsensiRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "format data tidak valid",
		})
		return
	}

	if err := services.ProsesUpdateAbsensi(uint(id), req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "Data Absensi berhasil diperbarui",
	})
}

func GetRekapAbsensi(ctx *gin.Context) {
	//AMbil id dari  TOKEN JWT
	userIDContext, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"succes":  false,
			"message": "Akses ditolak, token tidak valid",
		})
		return
	}

	userID := uint(userIDContext.(float64))
	//Panggil Service
	rekapData, err := services.GetRekapAbsensiMahasiswa(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "Berhasil mengambil rekap absensi",
		"data":    rekapData,
	})
}
