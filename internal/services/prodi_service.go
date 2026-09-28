package services

import (
	"errors"
	"siakad/internal/models"
	"siakad/internal/repositories"
)

func CreateProdi(kode, nama string) error {
	if kode == "" || nama == "" {
		return errors.New("Kode dan Nama Prodi tidak boleh kosong ")
	}
	prodi := &models.Prodi{
		Kode: kode,
		Nama: nama,
	}
	return repositories.CreateProdi(prodi)
}

func GetAllProdi() ([]models.Prodi, error) {
	return repositories.GetAllProdi()
}

func UpdateProdi(id uint, kode, nama string) error {
	prodi, err := repositories.GetProdiByID(id)
	if err != nil {
		return errors.New("Prodi tidak ditemukan")
	}
	prodi.Kode = kode
	prodi.Nama = nama
	return repositories.UpdateProdi(prodi)
}

func DeleteProdi(id uint) error {
	_, err := repositories.GetProdiByID(id)
	if err != nil {
		return errors.New("Prodi tidak ditemukan")
	}
	return repositories.DeleteProdi(id)
}
