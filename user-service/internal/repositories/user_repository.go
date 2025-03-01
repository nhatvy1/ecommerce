package repositories

import (
	"user-service/global"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserById() int
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		db: global.MySQL_Gorm,
	}
}

func (up *userRepository) GetUserById() int {
	return 1
}
