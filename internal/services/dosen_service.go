package services

import (
	"errors"
	"siakad/internal/models"
	"siakad/internal/repositories"
	"siakad/internal/utils"
)

func CreateDosen(username, password, nidn, nama string, prodiID uint) error {
	_, err := repositories.GetProdiByID(prodiID)
	if err != nil {
		return errors.New("Prodi tidak ditemukan")
	}

	role, err := repositories.GetOrCreateRole("dosen")
	if err != nil {
		return errors.New("gagal menyiapkan role dosen")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return errors.New("gagal mengenkripsi password")
	}

	user := &models.User{
		Username: username,
		Password: hashedPassword,
		RoleID:   role.ID,
	}

	dosen := &models.Dosen{
		NIDN:    nidn,
		Nama:    nama,
		ProdiID: prodiID,
	}

	return repositories.CreateDosenWithUser(user, dosen)
}

func GetAllDosen() ([]models.Dosen, error) {
	return repositories.GetAllDosen()
}

func UpdateDosen(id uint, nidn, nama string, prodiID uint) error {
	dosen, err := repositories.GetDosenByID(id)
	if err != nil {
		return errors.New("Data Dosen Tidak ditemukan")
	}

	_, errProdi := repositories.GetProdiByID(prodiID)
	if errProdi != nil {
		return errors.New("Prodi Tidak Valid")
	}

	dosen.NIDN = nidn
	dosen.Nama = nama
	dosen.ProdiID = prodiID

	return repositories.UpdateDosen(dosen)
}

func DeleteDosen(id uint) error {
	_, err := repositories.GetDosenByID(id)
	if err != nil {
		return errors.New("data dosen tidak ditemukan")
	}
	return repositories.DeleteDosen(id)
}
