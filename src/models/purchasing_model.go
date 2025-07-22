package models

type PurcahsingModel struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ProductID   int    `json:"product_id"`
	FactureCode string `json:"facture_code"`
	Pemeriksa   string `json:"pemeriksa"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
}

func (PurcahsingModel) TableName() string {
	return "purchasing"
}
