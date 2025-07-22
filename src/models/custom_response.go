package models

import "gorm.io/gorm"

type CustomResponse struct {
	gorm.Model

	Data     interface{} `json:"data"`
	Httpcode int         `json:httocde`
	Message  string      `json:data`
}
