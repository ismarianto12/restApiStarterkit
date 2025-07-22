package models

type BarangModel struct {
	Id         int     `json:"id"  gorm:"column:id;primaryKey"`
	Kode       string  `json:"kode" binding:"required" gorm:"column:kode"`
	Nama       string  `json:"nama" binding:"required" gorm:"column:nama"`
	KategoryId int     `json:"kategory_id" gorm:"column:kategory_id"`
	Deskripsi  string  `json:"deskripsi" gorm:"column:deskripsi"`
	Gambar     string  `json:"gambar" gorm:"column:gambar"`
	CreatedAt  string  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  string  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt  *string `json:"deleted_at" gorm:"column:deleted_at"`
}

func (BarangModel) TableName() string {
	return "barang"
}
