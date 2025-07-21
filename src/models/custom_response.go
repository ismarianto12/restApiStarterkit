package models

import "gorm.io/gorm"

type CustomResponse struct {
	gorm.Model

	data     string `json:data`
	httpcode string `json:httocde`
	message  string `json:data`
}
