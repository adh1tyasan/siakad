package main

import (
	"log"
	"os"

	"siakad/config"
	"siakad/internal/models"
	"siakad/internal/routes"
	"siakad/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.ConnectDB()

	// Auto Migrate SELURUH model yang sudah kita buat
	err := config.DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Prodi{},
		&models.TahunAkademik{},
		&models.Ruangan{},
		&models.MataKuliah{},
		&models.Dosen{},
		&models.Mahasiswa{},
		&models.Kelas{},
		&models.Jadwal{},
		&models.KRS{},
		&models.KRSDetail{},
		&models.Presensi{},
		&models.PresensiDetail{},
		&models.Nilai{},
		&models.Pengumuman{},
	)
	if err != nil {
		log.Fatal("Gagal melakukan migrasi database:", err)
	}

	// SEED DATA ADMIN (Dijalankan sekali jika tabel role kosong)
	var count int64
	config.DB.Model(&models.Role{}).Count(&count)
	if count == 0 {
		adminRole := models.Role{Name: "admin"}
		config.DB.Create(&adminRole)

		hashedPassword, _ := utils.HashPassword("admin123")
		adminUser := models.User{
			Username: "admin",
			Password: hashedPassword,
			RoleID:   adminRole.ID,
		}
		config.DB.Create(&adminUser)
		log.Println("Berhasil membuat seed data: username 'admin', password 'admin123'")
	}

	router := gin.Default()

	// Daftarkan semua route
	routes.SetupRoutes(router)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server berjalan di http://localhost:%s\n", port)
	router.Run(":" + port)
}
