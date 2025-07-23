package models

type UserMode struct {
	ID       uint   `gorm:"primaryKey",json:"id"`
	Name     string `gorm:"not null",json:"name"`
	Username string `gorm:"not null",json:"username"`
	Password string `gorm:"not null",json:"password"`
	Image    string `gorm:"not null",json:"image"`
	Email    string `gorm:"unique;not null",json:"email"`
}

// TableName returns the name of the table in the database
func (UserMode) TableName() string {
	return "users"
}
