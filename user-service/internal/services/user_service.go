package services

import "user-service/internal/repositories"

type IUserService interface {
	GetUserById() int
}

type userService struct {
	userRepo repositories.IUserRepository
}

func NewUserService(userRepo repositories.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) GetUserById() int {
	return 10
}
