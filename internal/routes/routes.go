package routes

import (
	"siakad/internal/handlers"
	"siakad/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	//Endpoint public (Tanpa login)
	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	//Endpoint Terproteksi (Wajib Login JWT)
	protected := api.Group("/")
	protected.Use(middleware.RequireAuth())
	{
		//Contoh endpoint yang bisa diakses siapa saja yang sudah login
		protected.GET("/profile", func(ctx *gin.Context) {
			userID, _ := ctx.Get("user_id")
			role, _ := ctx.Get("role")
			ctx.JSON(200, gin.H{
				"succes":  true,
				"user_id": userID,
				"role":    role,
			})
		})
		protected.POST("/krs/tambah", handlers.TambahMataKuliahKRS)
		//Contoh endpoint yang bisa diakses oleh ADMIN
		adminOnly := protected.Group("/admin")
		adminOnly.Use(middleware.RequireRole("admin"))
		{
			adminOnly.GET("/dashboard", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"succes":  true,
					"message": "Selamat datang di Dashboard Admin",
				})
			})
			adminOnly.POST("/prodi", handlers.CreateProdi)
			adminOnly.GET("/prodi", handlers.GetAllProdi)
			adminOnly.PUT("/prodi/:id", handlers.UpdateProdi)
			adminOnly.DELETE("/prodi/:id", handlers.DeleteProdi)
			//CRUD Mahasiswa
			adminOnly.POST("/mahasiswa", handlers.CreateMahasiswa)
			adminOnly.GET("/mahasiswa", handlers.GetAllMahasiswa)
			adminOnly.PUT("/mahasiswa/:id", handlers.UpdateMahasiswa)
			adminOnly.DELETE("/mahasiswa/:id", handlers.DeleteMahasiswa)
			//CRUD Dosen
			adminOnly.POST("/dosen", handlers.CreateDosen)
			adminOnly.GET("/dosen", handlers.GetAllDosen)
			adminOnly.PUT("/dosen/:id", handlers.UpdateDosen)
			adminOnly.DELETE("/dosen/:id", handlers.DeleteDosen)
			//CRUD MataKuliah
			adminOnly.POST("/mata-kuliah", handlers.CreataMataKuliah)
			adminOnly.GET("/mata-kuliah", handlers.GetAllMataKuliah)
			adminOnly.PUT("/mata-kuliah/:id", handlers.UpdateMataKuliah)
			adminOnly.DELETE("/mata-kuliah/:id", handlers.DeleteMataKuliah)
			//CRUD JADWAL
			adminOnly.POST("/jadwal", handlers.CreateJadwal)
			adminOnly.GET("/jadwal", handlers.GetAllJadwal)
			adminOnly.PUT("/jadwal/:id", handlers.UpdateJadwal)
			adminOnly.DELETE("/jadwal/:id", handlers.DeleteJadwal)
			//CRUD KELAS
			adminOnly.POST("/kelas", handlers.CreateKelas)
			adminOnly.GET("/kelas", handlers.GetAllKelas)
			adminOnly.PUT("/kelas/:id", handlers.UpdateKelas)
			adminOnly.DELETE("kelas/:id", handlers.DeleteKelas)
		}
		mahasiswaOnly := protected.Group("/mahasiswa")
		mahasiswaOnly.Use(middleware.RequireRole("mahasiswa"))
		{
			mahasiswaOnly.POST("/krs/:id/submit", handlers.SubmitKRS)
		}
		dosenOnly := protected.Group("/dosen")
		dosenOnly.Use(middleware.RequireRole("dosen"))
		{
			dosenOnly.PUT("/krs/:id/review-pa", handlers.ReviewKRSByPA)
		}
		kaprodiOnly := protected.Group("/kaprodi")
		kaprodiOnly.Use(middleware.RequireRole("kaprodi", "admin"))
		{
			kaprodiOnly.PUT("/krs/:id/review-kaprodi", handlers.ReviewKRSByKaprodi)
		}
	}
}
