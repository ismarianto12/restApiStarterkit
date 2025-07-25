package repository

import (
	"golangRest/src/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (ptk *CategoryRepository) GetAll() (*models.CategoryModel, error) {

	var models models.CategoryModel

	if err := ptk.DB.Table("category").Find(&models).Scan(&models).Error; err != nil {
		return nil, err
	}
	return &models, nil
}
