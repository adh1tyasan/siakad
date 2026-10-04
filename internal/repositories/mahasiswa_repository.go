package repositories

import (
	"siakad/config"
	"siakad/internal/models"

	"gorm.io/gorm"
)

// GetAndCreateRole memastikan role tersedia di database
func GetOrCreateRole(roleName string) (*models.Role, error) {
	var role models.Role
	err := config.DB.FirstOrCreate(&role, models.Role{Name: roleName}).Error
	return &role, err
}

// CreateMahasiswaWithUser menggunakan transaction. Jika salah satu gagal, maka semua dibatalkan
func CreateMahasiswaWithUser(user *models.User, mhs *models.Mahasiswa) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		//Hubungkan ID User yang baru dibuat ke profil Mahasiswa
		mhs.UserID = user.ID
		if err := tx.Create(mhs).Error; err != nil {
			return err
		}
		return nil
	})
}

func GetAllMahasiswa() ([]models.Mahasiswa, error) {
	var mahasiswas []models.Mahasiswa
	//Preload untuk menarik data user dan prodi sekaligus (JOIN)
	err := config.DB.Preload("User").Preload("Prodi").Find(&mahasiswas).Error
	return mahasiswas, err
}

func GetMahasiswaByID(id uint) (*models.Mahasiswa, error) {
	var mhs models.Mahasiswa
	err := config.DB.Preload("User").Preload("Prodi").First(&mhs, id).Error
	return &mhs, err
}

func UpdateMahasiswa(mhs *models.Mahasiswa) error {
	return config.DB.Save(mhs).Error
}

func DeleteMahasiswa(id uint) error {
	//Menghapus data mahasiswa. GORM otomatis melakukan soft delete
	return config.DB.Delete(&models.Mahasiswa{}, id).Error
}

func GetMahasiswaByUserID(userID uint) (*models.Mahasiswa, error) {
	var mhs models.Mahasiswa
	err := config.DB.Where("user_id = ?", userID).First(&mhs).Error
	return &mhs, err
}
