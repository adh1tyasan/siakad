package handlers

import (
	"net/http"
	"siakad/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewKRSRequest struct {
	Status  string `json:"status" binding:"required"`
	Catatan string `json:"catatan"`
}

// Handler Untuk mahasiswa
func SubmitKRS(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := services.AjukanKRS(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "KRS Berhasil diajukan ke Dosen PA",
	})
}

// Handler untuk Dosen PA
func ReviewKRSByPA(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req ReviewKRSRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "Format data salah",
		})
		return
	}
	if err := services.PersetujuanPA(uint(id), req.Status, req.Catatan); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "Berhasil memproses KRS di tingkat PA",
	})
}

// Handler untuk Kaprodi
func ReviewKRSByKaprodi(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req ReviewKRSRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "format data salah",
		})
		return
	}
	if err := services.PersetujuanKaprodi(uint(id), req.Status, req.Catatan); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "Berhasil Memproses KRS di tingkat Kaprodi",
	})
}

func TambahMataKuliahKRS(ctx *gin.Context) {
	var req services.TambahKRSRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "Format data tidak valid",
		})
		return
	}
	//Panggil service untuk logika penambahan KRS
	err := services.TambahJadwalKeKRS(req)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"succes":  true,
		"message": "Jadwal Berhasil ditambahkan ke KRS",
	})
}
