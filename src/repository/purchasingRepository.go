package repository

import (
	"golangRest/src/models"

	"gorm.io/gorm"
)

type PurcashingRepository struct {
	DB *gorm.DB
}

func NewPurcashingRepository(db *gorm.DB) *PurcashingRepository {
	return &PurcashingRepository{DB: db}
}
func (rpcp *PurcashingRepository) GetAll() ([]models.PurcahsingModel, error) {
	var purchasing []models.PurcahsingModel
	if err := rpcp.DB.Find(&purchasing).Error; err != nil {
		return nil, err
	}
	return purchasing, nil
}
func (rpc *PurcashingRepository) Create(purchasing *models.PurcahsingModel) error {
	if err := rpc.DB.Create(&purchasing).Error; err != nil {
		return err
	}
	return nil
}
func (rpc *PurcashingRepository) upDateData(id uint) (purchasing *models.PurcahsingModel) {
	rpc.DB.Find(id).Save(purchasing)
	return
}
