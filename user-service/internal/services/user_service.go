package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"user-service/internal/database"
	"user-service/internal/model"
	"user-service/internal/repositories"
	"user-service/internal/vo"
	"user-service/pkg/response"
	"user-service/pkg/utils/auth"
	copy_helper "user-service/pkg/utils/copyhelper"
	"user-service/pkg/utils/crypto"
	utils "user-service/pkg/utils/uuid"
)

type (
	IUserService interface {
		// GetUserById(ctx context.Context, id int) (*database.User, error)
		Register(ctx context.Context, user *vo.UserRegister) (int, error)
		UpdateUser(ctx context.Context) int
		Login(ctx context.Context, user *vo.UserLogin) (int, model.LoginOutput, error)
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

func (us *userService) Register(ctx context.Context, userRegister *vo.UserRegister) (int, error) {
	userFound, err := us.userRepo.CheckUserEmail(ctx, userRegister.UserEmail)

	if err != nil {
		return response.ErrCodeUserHasExists, err
	}

	if userFound > 0 {
		return response.ErrCodeUserHasExists, fmt.Errorf("user has already registered")
	}

	user := database.InsertUserBaseParams{}
	userSalt, err := crypto.GenerateSalt(16)
	if err != nil {
		return response.ErrPlsTryAgainLater, err
	}
	userPassword := crypto.HashPassword(userRegister.UserPassword, userSalt)
	us.modelConverter.Copy(&user, &userRegister)
	user.UserSalt = userSalt
	user.UserPassword = userPassword

	newUserBase, err := us.userRepo.Create(ctx, &user)
	if err != nil {
		return response.ErrPlsTryAgainLater, err
	}

	return newUserBase, nil
}

func (us *userService) Login(ctx context.Context, userLogin *vo.UserLogin) (codeResult int, out model.LoginOutput, err error) {
	userBase, err := us.userRepo.FindUserByEmail(ctx, userLogin.UserEmail)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeAuthFailed, out, fmt.Errorf("email or password incorrect")
		}
		return response.ErrCodeAuthFailed, out, err
	}

	if !crypto.MatchingPassword(userBase.UserPassword, userLogin.UserPassword, userBase.UserSalt) {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("email or password incorrect")
	}

	subToken := utils.GenerateCliTokenUUID(int(userBase.UserID))
	out.Token, err = auth.CreateToken(subToken)
	if err != nil {
		return
	}

	return 200, out, nil
}

func (us *userService) UpdateUser(ctx context.Context) int {

	data := us.userRepo.Update(ctx)
	return data
}
