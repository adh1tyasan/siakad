package services

import (
	"errors"
	"siakad/internal/models"
	"siakad/internal/repositories"
	"siakad/internal/utils"
)

func CreateMahasiswa(username, password, npm, nama, kelas string, prodiID uint) error {
	//Validasi apakah ProdiID yang dimasukkan ada di tabel prodi
	_, err := repositories.GetProdiByID(prodiID)
	if err != nil {
		return errors.New("Prodi Tidak ditemukan")
	}

	//Ambil atau buat role "Mahasiswa"
	role, err := repositories.GetOrCreateRole("mahasiswa")
	if err != nil {
		return errors.New("gagal menyiapkan role mahasiswa")
	}

	//Enkripsi password sebelum masuk kedatabase
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return errors.New("Gagal mengenkripsi password")
	}

	//Kerangka data
	user := &models.User{
		Username: username,
		Password: hashedPassword,
		RoleID:   role.ID,
	}

	mahasiswa := &models.Mahasiswa{
		NPM:     npm,
		Nama:    nama,
		ProdiID: prodiID,
		Kelas:   kelas,
	}

	//Simpan ke database
	return repositories.CreateMahasiswaWithUser(user, mahasiswa)
}

func GetAllMahasiswa() ([]models.Mahasiswa, error) {
	return repositories.GetAllMahasiswa()
}

func UpdateMahasiswa(id uint, npm, nama, kelas string, prodiID uint) error {
	mhs, err := repositories.GetMahasiswaByID(id)
	if err != nil {
		return errors.New("Data Mahasiswa Tidak ditemukan")
	}

	_, errProdi := repositories.GetProdiByID(prodiID)
	if errProdi != nil {
		return errors.New("Prodi Tidak Valid")
	}

	mhs.NPM = npm
	mhs.Nama = nama
	mhs.Kelas = kelas
	mhs.ProdiID = prodiID

	return repositories.UpdateMahasiswa(mhs)
}

func DeleteMahasiswa(id uint) error {
	_, err := repositories.GetMahasiswaByID(id)
	if err != nil {
		return errors.New("Data Mahasiswa Tidak ditemukan")
	}

	return repositories.DeleteMahasiswa(id)
}
