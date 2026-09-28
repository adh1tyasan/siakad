package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL belum di set di .env")
	}

	// Tambahkan konfigurasi NamingStrategy di sini
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // GORM tidak akan lagi menambahkan "s" di akhir nama tabel
		},
	})
	if err != nil {
		log.Fatal("Gagal terhubung ke database", err)
	}

	DB = database
	log.Println("Database supabase Berhasil terhubung")
}
