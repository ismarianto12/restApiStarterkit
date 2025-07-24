package repository

import (
	"golangRest/src/models"

	"gorm.io/gorm"
)

type SuplierRepository struct {
	DB *gorm.DB
}

func NewSuplierRepositoryInstance(db *gorm.DB) *SuplierRepository {
	return &SuplierRepository{DB: db}
}

func (rp *SuplierRepository) GetallData() ([]map[string]interface{}, error) {
	var spr []map[string]interface{}
	if err := rp.DB.Raw("select * from suplier").Scan(&spr).Error; err != nil {
		return nil, err
	}
	return spr, nil
}

func (rp *SuplierRepository) InserDatat(md *models.SuplierModel) error {
	if err := rp.DB.Save(md).Error; err != nil {
		return nil
	}
	return nil
}

func (rp *SuplierRepository) UpdateData(id int, md *models.SuplierModel) error {

	var existingdata models.SuplierModel
	if err := rp.DB.First(id).Error; err != nil {
		return nil
	}
	existingdata.Kode = md.Kode
	existingdata.Nama = md.Nama
	existingdata.Address = md.Address

	if err := rp.DB.Save(existingdata).Error; err != nil {
		return nil
	}
	return nil
}

func (rp *SuplierRepository) DeleteData(id int, md *models.SuplierModel) error {
	if err := rp.DB.Find(id).Delete(md).Error; err != nil {
		return nil
	}
	return nil
}
