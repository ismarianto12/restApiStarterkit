package models

import (
	"gorm.io/gorm"
)

type BarangModel struct {
	gorm.Model
	Id         string `gorm:"primaryKey",json:"id`
	Kode       string `gorm:"primaryKey",json:"kode`
	Nama       string `gorm:"primaryKey",json:"nama`
	KategoriId string `gorm:"primaryKey",json:"kategory_id`
	Deskripsi  string `gorm:"primaryKey",json:"deskripsi`
	Gambar     string `gorm:"primaryKey",json:"gambar`
	created_at string `gorm:"primaryKey",json:"createdAt`
	updated_at string `gorm:"primaryKey",json:"updatedAt`
}

func (BarangModel) TableName() string {
	return "barang"
}
