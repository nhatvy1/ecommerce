package repositories

import (
	"user-service/global"
	"user-service/internal/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserById(id int) (*models.APIUser, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		db: global.MySQL_Gorm,
	}
}

func (up *userRepository) GetUserById(id int) (*models.APIUser, error) {
	user := models.APIUser{}
	if err := up.db.Model(&models.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
