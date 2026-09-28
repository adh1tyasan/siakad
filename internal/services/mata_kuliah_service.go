package services

import (
	"errors"
	"siakad/internal/models"
	"siakad/internal/repositories"
)

func CreateMataKuliah(kode, nama string, sks, semester int, prodiID uint) error {
	//Validasi apakah prodi yang di input ada atau tidak
	_, err := repositories.GetProdiByID(prodiID)
	if err != nil {
		return errors.New("Prodi tidak valid atau tidak ditemukan")
	}

	mk := &models.MataKuliah{
		Kode:     kode,
		Nama:     nama,
		SKS:      sks,
		Semester: semester,
		ProdiID:  prodiID,
	}

	return repositories.CreateMataKuliah(mk)
}

func GetAllMataKuliah() ([]models.MataKuliah, error) {
	return repositories.GetAllMataKuliah()
}

func UpdateMataKuliah(id uint, kode, nama string, sks, semester int, prodiID uint) error {
	mk, err := repositories.GetMataKuliahByID(id)
	if err != nil {
		return errors.New("Data Mata Kuliah tidak ditemukan")
	}

	_, errProdi := repositories.GetProdiByID(prodiID)
	if errProdi != nil {
		return errors.New("Prodi tidak valid")
	}

	mk.Kode = kode
	mk.Nama = nama
	mk.SKS = sks
	mk.Semester = semester
	mk.ProdiID = prodiID

	return repositories.UpdateMataKuliah(mk)
}

func DeleteMataKuliah(id uint) error {
	_, err := repositories.GetMataKuliahByID(id)
	if err != nil {
		return errors.New("data mata kuliah tidak dapat ditemukan")
	}
	return repositories.DeleteMataKuliah(id)
}
