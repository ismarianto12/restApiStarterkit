package models

type SuplierModel struct {
	id         int    `json:"id"`
	Nama       string `json:"nama"`
	Kode       string `json:"nama"`
	Contact    string `json:"nama"`
	Address    string `json:"nama"`
	created_at string `json:"nama"`
}

func (SuplierModel) TableName() string {
	return "suplier"
}
