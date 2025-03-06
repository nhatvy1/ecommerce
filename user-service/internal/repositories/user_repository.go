package repositories

import (
	"user-service/global"
	"user-service/internal/models"

	"gorm.io/gorm"
)

type (
	IUserRepository interface {
		FindById(id int) (*models.APIUser, error)
		FindByEmail(email string) (*models.APIUserEmail, error)
		Create(user *models.User) (*models.User, error)
		Update(id int, user *models.User) (*models.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository() IUserRepository {
	return &userRepository{
		db: global.MySQL_Gorm,
	}
}

func (userRepo *userRepository) FindById(id int) (*models.APIUser, error) {
	user := models.APIUser{}
	if err := userRepo.db.Model(&models.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepo *userRepository) FindByEmail(email string) (*models.APIUserEmail, error) {
	user := models.APIUserEmail{}

	if err := userRepo.db.Model(&models.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepo *userRepository) Create(user *models.User) (*models.User, error) {
	if err := userRepo.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (userRepo *userRepository) Update(id int, user *models.User) (*models.User, error) {
	if err := userRepo.db.Where("id = ?", id).Updates(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
