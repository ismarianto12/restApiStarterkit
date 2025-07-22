package repository

import (
	"fmt"
	"golangRest/src/models"

	"gorm.io/gorm"
)

type BarangRepository struct {
	DB *gorm.DB
}

func NewBarangRepository(db *gorm.DB) *BarangRepository {
	return &BarangRepository{DB: db}
}

func (repo *BarangRepository) GetAllBarang(page int) (*models.BarangModel, error) {
	fmt.Println(`$page`)
	var barangModel models.BarangModel
	if err := repo.DB.Exec("select * from barang").Error; err != nil {
		return nil, err
	}
	return &barangModel, nil
}
func (repo *BarangRepository) InsertData(barang *models.BarangModel) error {
	fmt.Println("insert")
	if err := repo.DB.Save(barang).Error; err != nil {
		return err
	}
	return nil
}

func (repo *BarangRepository) Update(id int, newData *models.BarangModel) (*models.BarangModel, error) {
	var barang models.BarangModel

	// Cari data lama berdasarkan ID
	if err := repo.DB.First(&barang, id).Error; err != nil {
		return nil, err // data tidak ditemukan
	}

	// Update field dengan data baru
	barang.Nama = newData.Nama
	barang.Gambar = newData.Gambar
	barang.Kode = newData.Kode
	if err := repo.DB.Save(&barang).Error; err != nil {
		return nil, err
	}

	return &barang, nil
}
