package handlers

import (
	"net/http"
	"siakad/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JadwalInput struct {
	TahunAkademikID uint   `json:"tahun_akademik_id" binding:"required"`
	MataKuliahID    uint   `json:"mata_kuliah_id" binding:"required"`
	KelasID         uint   `json:"kelas_id" binding:"required"`
	DosenID         uint   `json:"dosen_id" binding:"required"`
	RuanganID       uint   `json:"ruangan_id" binding:"required"`
	Hari            string `json:"hari" binding:"required"`
	JamMulai        string `json:"jam_mulai" binding:"required"`
	JamSelesai      string `json:"jam_selesai" binding:"required"`
}

func CreateJadwal(ctx *gin.Context) {
	var input JadwalInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "Format input salah",
		})
		return
	}

	err := services.CreateJadwal(input.TahunAkademikID, input.MataKuliahID, input.KelasID, input.DosenID, input.RuanganID, input.Hari, input.JamMulai, input.JamSelesai)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"succes":  true,
		"message": "Jadwal Berhasil ditambahkan",
	})
}

func GetAllJadwal(ctx *gin.Context) {
	jadwals, err := services.GetAllJadwal()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"succes":  false,
			"message": "gagal mengambil data jadwal",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes": true,
		"data":   jadwals,
	})
}

func UpdateJadwal(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input JadwalInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":  false,
			"message": "format input salah",
		})
		return
	}

	err := services.UpdateJadwal(uint(id), input.TahunAkademikID, input.MataKuliahID, input.KelasID, input.DosenID, input.RuanganID, input.Hari, input.JamMulai, input.JamSelesai)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "Data Jadwal berhasil diperbarui",
	})
}

func DeleteJadwal(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := services.DeleteJadwal(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"succes":  true,
		"message": "Data jadwal berhasil dihapus",
	})
}
