package services

import (
	"errors"
	"siakad/internal/models"
	"siakad/internal/repositories"
)

func CreateKelas(nama string, prodiID uint, tahunAkademikID uint) error {
	if nama == "" {
		return errors.New("Nama Kelas tidak boleh kosong")
	}

	kelas := &models.Kelas{
		Nama:            nama,
		ProdiID:         prodiID,
		TahunAkademikID: tahunAkademikID,
	}

	return repositories.CreateKelas(kelas)
}

func GetAllKelas() ([]models.Kelas, error) {
	return repositories.GetAllKelas()
}

func UpdateKelas(id uint, nama string, prodiID uint, tahunAkademikID uint) error {
	kelas, err := repositories.GetKelasByID(id)
	if err != nil {
		return errors.New("Kelas tidak ditemukan")
	}
	kelas.Nama = nama
	kelas.ProdiID = prodiID
	kelas.TahunAkademikID = tahunAkademikID

	return repositories.UpdateKelas(kelas)
}

func DeleteKelas(id uint) error {
	_, err := repositories.GetKelasByID(id)
	if err != nil {
		return errors.New("Kelas tidak ditemukan")
	}
	return repositories.DeleteKelas(id)
}
