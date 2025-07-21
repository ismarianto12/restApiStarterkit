package models

import "gorm.io/gorm"

type PurcahsingModel struct {
	gorm.Model
	ProductID   string `gorm:"primaryKey",json:product_id`
	FactureCode string `gorm:"not null",json:facture_code`
	Pemeriksa   string `json:pemeriksa`
	CreatedBy   string `json:created_by`
	UpdatedBy   string `json:updated_by`
}

func (PurcahsingModel) TableName() string {
	return "purchasing"
}
