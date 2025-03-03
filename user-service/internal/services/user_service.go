package services

import (
	"user-service/internal/models"
	"user-service/internal/repositories"
)

type IUserService interface {
	GetUserById(id int) (*models.APIUser, error)
	Register() (int, error)
}

type userService struct {
	userRepo repositories.IUserRepository
}

func NewUserService(userRepo repositories.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) GetUserById(id int) (*models.APIUser, error) {
	user, err := us.userRepo.GetUserById(id)
	return user, err
}

func (us *userService) Register() (int, error) {
	return 1, nil
}
