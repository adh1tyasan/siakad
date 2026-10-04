package repositories

import (
	"siakad/config"
	"siakad/internal/models"
)

// Mengambil Jadwal beserta informasi SKS dari Mata Kuliah
func GetJadwalBerdasarkanID(jadwalID uint) (*models.Jadwal, error) {
	var jadwal models.Jadwal
	err := config.DB.Preload("MataKuliah").Where("id = ?", jadwalID).First(&jadwal).Error
	return &jadwal, err
}

// Mencari KRS Mahasiswa pada semester tertentu yang statusnya masih draft
func GetDraftKRSMahasiswa(mahasiswaID, tahunAkademikID uint) (*models.KRS, error) {
	var krs models.KRS
	err := config.DB.Preload("KRSDetails.Jadwal").Where("mahasiswa_id = ? AND tahun_akademik_id = ? AND status = ?", mahasiswaID, tahunAkademikID, "Draft").First(&krs).Error
	return &krs, err
}

// Membuat KRS baru jika mahasiswa belum punya KRS di semester ini
func CreateKRSBaru(krs *models.KRS) error {
	return config.DB.Create(&krs).Error
}

// Simpan DetailKrs (jadwal yang dipilih) dan update total sks
func TambahDetailKRS(krsID uint, detail *models.KRSDetail, totalSKSBaru int) error {
	tx := config.DB.Begin()

	//SImpan detail(jadwal yang diambil)
	if err := tx.Create(detail).Error; err != nil {
		tx.Rollback()
		return err
	}

	//Update total SKS di tabel induk KRS
	if err := tx.Model(&models.KRS{}).Where("id = ?", krsID).Updates(map[string]any{
		"total_sks": totalSKSBaru,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func GetKRSByID(id uint) (*models.KRS, error) {
	var krs models.KRS
	err := config.DB.Preload("Mahasiswa").First(&krs, id).Error
	return &krs, err
}

func UpdateStatusKRS(id uint, status, catatan string) error {
	return config.DB.Model(&models.KRS{}).Where("id = ?", id).Updates(map[string]any{
		"status":  status,
		"catatan": catatan,
	}).Error
}

func GetKRSByMahasiswaID(mahasiswaID uint) ([]models.KRS, error) {
	var krs []models.KRS

	err := config.DB.
		Preload("TahunAkademik").
		Preload("KRSDetail").
		Preload("KRSDetail.Jadwal").
		Preload("KRSDetail.Jadwal.MataKuliah").
		Preload("KRSDetail.Jadwal.Dosen").
		Preload("KRSDetail.Jadwal.Ruangan").
		Where("mahasiswa_id = ?", mahasiswaID).Find(&krs).Error
	return krs, err

}
