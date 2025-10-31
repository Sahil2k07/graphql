package repositories

import (
	"github.com/Sahil2k07/graphql/internal/database"
	"github.com/Sahil2k07/graphql/internal/interfaces"
	"github.com/Sahil2k07/graphql/internal/models"
)

type authRepository struct{}

func NewAuthRepository() interfaces.AuthRepository {
	return &authRepository{}
}

func (r *authRepository) CheckUserExist(email string) (bool, error) {
	var count int64

	err := database.DB.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return true, err
	}

	return count > 0, nil
}

func (r *authRepository) GetUser(email string) (models.User, error) {
	var user models.User

	err := database.DB.Preload("Profile").Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *authRepository) AddUser(user models.User) error {
	return database.DB.Create(&user).Error
}

func (r *authRepository) UpdatePassword(email, newPassword string) error {
	return database.DB.Model(&models.User{}).Where("email = ?", email).Update("password", newPassword).Error
}
