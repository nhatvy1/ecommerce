package services

import (
	"errors"
	"user-service/internal/models"
	"user-service/internal/repositories"
	"user-service/internal/vo"

	"github.com/jinzhu/copier"
)

type IUserService interface {
	GetUserById(id int) (*models.APIUser, error)
	Register(user *vo.UserRegister) (int, error)
	IsEmailRegistered(email string) (bool, error)
	UpdateUser(id int, user *vo.UserUpdate) (int, error)
}

type userService struct {
	userRepo repositories.IUserRepository
}

func NewUserService(userRepo repositories.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) IsEmailRegistered(email string) (bool, error) {
	userFound, err := us.userRepo.GetUserByEmail(email)
	if err != nil {
		return false, err
	}

	if userFound != nil {
		return true, nil
	}

	return false, nil
}

func (us *userService) GetUserById(id int) (*models.APIUser, error) {
	user, err := us.userRepo.GetUserById(id)
	return user, err
}

func (us *userService) Register(userRegister *vo.UserRegister) (int, error) {
	userFound, _ := us.IsEmailRegistered(userRegister.Email)

	if userFound {
		return -1, errors.New("email already exists")
	}

	user := models.User{}
	copier.Copy(&user, &userRegister)

	_, err := us.userRepo.Create(&user)
	if err != nil {
		return -1, err
	}

	return 1, nil
}

func (us *userService) UpdateUser(id int, userUpdate *vo.UserUpdate) (int, error) {
	_, errUserNotFound := us.userRepo.GetUserById(id)
	if errUserNotFound != nil {
		return -1, errors.New("user not found")
	}

	user := models.User{}
	copier.Copy(&user, &userUpdate)
	if userUpdate.Status != nil {
		user.Status = models.ToStatus(*userUpdate.Status)
	}

	_, err := us.userRepo.UpdateUser(id, &user)
	if err != nil {
		return -1, errors.New("user not found")
	}

	return 1, nil
}
