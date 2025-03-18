package repositories

import (
	"context"
	"user-service/global"
	"user-service/internal/database"
	"user-service/pkg/response"
)

type (
	IUserRepository interface {
		// FindById(ctx context.Context, id int) (*database.User, error)
		Create(ctx context.Context, user *database.InsertUserBaseParams) (int, error)
		Update(ctx context.Context) int
		CheckUserEmail(ctx context.Context, email string) (int64, error)
		FindUserByEmail(ctx context.Context, email string) (*database.UserBase, error)
	}

	userRepository struct {
		db *database.Queries
	}
)

func NewUserRepository() IUserRepository {
	return &userRepository{
		db: database.New(global.MySQL_SQLC),
	}
}

func (userRpo *userRepository) FindUserByEmail(ctx context.Context, email string) (*database.UserBase, error) {
	data, err := userRpo.db.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (userRepo *userRepository) CheckUserEmail(ctx context.Context, email string) (int64, error) {
	data, err := userRepo.db.CheckUserExists(ctx, email)

	if err != nil {
		return -1, err
	}

	return data, nil
}

func (userRepo *userRepository) Create(ctx context.Context, user *database.InsertUserBaseParams) (int, error) {

	_, err := userRepo.db.InsertUserBase(ctx, *user)

	if err != nil {
		return response.ErrPlsTryAgainLater, err
	}

	return response.ErrCodeSuccess, nil
}

func (userRepo *userRepository) Update(ctx context.Context) int {
	return 1
}
