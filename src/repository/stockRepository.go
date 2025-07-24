package repository

import (
	"golangRest/src/models"

	"gorm.io/gorm"
)

type StockRepository struct {
	DB *gorm.DB
}

func NewStockRepository(db *gorm.DB) *StockRepository {
	return &StockRepository{DB: db}
}

func (rp *StockRepository) GetAllData() ([]models.StockModel, error) {
	var stock []models.StockModel

	if err := rp.DB.Find(&stock).Error; err != nil {
		return nil, err
	}
	return stock, nil
}

func (rp *StockRepository) FindById(id int) {
	var stock models.StockModel
	rp.DB.Find(&stock, id).Scan(&stock)
}

func (rp *StockRepository) DeleteById(id int) {
	var stock models.StockModel
	rp.DB.Delete(&stock, id)
}

func (rp *StockRepository) Update(id int, payloadupdate *models.StockModel) (*models.StockModel, error) {
	var stock models.StockModel
	if err := rp.DB.Find(&stock, id).Scan(&stock).Error; err != nil {
		return nil, err
	}
	stock.BarangId = payloadupdate.BarangId
	stock.BarangId = payloadupdate.BarangId

	if err := rp.DB.Save(payloadupdate).Error; err != nil {
		return nil, err
	}
	return &stock, nil
}
