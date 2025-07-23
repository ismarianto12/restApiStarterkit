package models

type StockModel struct {
	Id        int     `json:"id"  gorm:"column:id;primaryKey"`
	BarangId  string  `json:"barang_id" binding:"required" gorm:"column:barang_id"`
	Quantity  int     `json:"nama" binding:"required" gorm:"column:nama"`
	Deskripsi string  `json:"deskripsi" gorm:"column:deskripsi"`
	CreatedAt string  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *string `json:"deleted_at" gorm:"column:deleted_at"`
}

func (StockModel) TableName() string {
	return "stock"
}
