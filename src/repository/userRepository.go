package repository

import (
	"golangRest/src/database"
	"golangRest/src/models"

	"gorm.io/gorm"
)

// UserRepository provides methods to interact with the user database
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.DB,
	}
}

// GetUser retrieves a user by ID
func (repo *UserRepository) GetUser(id uint) (*models.UserMode, error) {
	var user models.UserMode
	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user in the database
func (repo *UserRepository) CreateUser(user *models.UserMode) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser updates an existing user in the database
func (repo *UserRepository) UpdateUser(user *models.UserMode) error {
	if err := repo.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user from the database
func (repo *UserRepository) DeleteUser(id uint) error {
	if err := repo.db.Delete(&models.UserMode{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetAllUsers() ([]models.UserMode, error) {
	var users []models.UserMode
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByEmail retrieves a user by email
func (repo *UserRepository) GetUserByEmail(email string) (*models.UserMode,
	error) {
	var user models.UserMode
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
