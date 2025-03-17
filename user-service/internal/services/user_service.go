package services

import (
	"context"
	"errors"
	"user-service/internal/database"
	"user-service/internal/repositories"
	"user-service/internal/vo"
	copy_helper "user-service/pkg/utils/copyhelper"

	"github.com/jinzhu/copier"
)

type (
	IUserService interface {
		GetUserById(ctx context.Context, id int) (*database.User, error)
		Register(ctx context.Context, user *vo.UserRegister) (int, error)
		IsEmailRegistered(ctx context.Context, email string) (bool, error)
		UpdateUser(ctx context.Context) int
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

func (us *userService) IsEmailRegistered(ctx context.Context, email string) (bool, error) {
	userFound, err := us.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if userFound != nil {
		return true, nil
	}

	return false, nil
}

func (us *userService) GetUserById(ctx context.Context, id int) (*database.User, error) {
	user, err := us.userRepo.FindById(ctx, id)
	return user, err
}

func (us *userService) Register(ctx context.Context, userRegister *vo.UserRegister) (int, error) {
	userFound, _ := us.IsEmailRegistered(ctx, userRegister.Email)

	if userFound {
		return -1, errors.New("email already exists")
	}

	user := database.User{}
	copier.Copy(&user, &userRegister)

	_, err := us.userRepo.Create(ctx, &user)
	if err != nil {
		return -1, err
	}

	return 1, nil
}

func (us *userService) UpdateUser(ctx context.Context) int {

	data := us.userRepo.Update(ctx)
	return data
}
