package models

type CategoryModel struct {
	id           int    `json:"id"`
	categoryName string `json:"string"`
	categoryCode string `json:string`
	created_at   string `json:created_at`
	updated_at   string `json:updated_at`
}

func TableName(CategoryModel) string {
	return "category"
}
