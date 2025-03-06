package services

import (
	"errors"
	"user-service/internal/models"
	"user-service/internal/repositories"
	"user-service/internal/vo"
	copy_helper "user-service/pkg/utils/copyhelper"

	"github.com/jinzhu/copier"
)

type (
	IUserService interface {
		GetUserById(id int) (*models.APIUser, error)
		Register(user *vo.UserRegister) (int, error)
		IsEmailRegistered(email string) (bool, error)
		UpdateUser(id int, user *vo.UserUpdate) (int, error)
	}

	userService struct {
		userRepo       repositories.IUserRepository
		modelConverter copy_helper.ModelConverter
	}
)

func NewUserService(userRepo repositories.IUserRepository, modelConverter copy_helper.ModelConverter) IUserService {
	return &userService{
		userRepo:       userRepo,
		modelConverter: modelConverter,
	}
}

func (us *userService) IsEmailRegistered(email string) (bool, error) {
	userFound, err := us.userRepo.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if userFound != nil {
		return true, nil
	}

	return false, nil
}

func (us *userService) GetUserById(id int) (*models.APIUser, error) {
	user, err := us.userRepo.FindById(id)
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
	_, errUserNotFound := us.userRepo.FindById(id)
	if errUserNotFound != nil {
		return -1, errors.New("user not found")
	}

	user := models.User{}

	us.modelConverter.Copy(&user, &userUpdate)

	if userUpdate.Status != nil {
		user.Status = models.ToStatus(*userUpdate.Status)
	}

	_, err := us.userRepo.Update(id, &user)
	if err != nil {
		return -1, errors.New("user not found")
	}

	return 1, nil
}
