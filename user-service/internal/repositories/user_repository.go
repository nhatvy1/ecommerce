package repositories

import (
	"context"
	"user-service/global"
	"user-service/internal/database"
)

type (
	IUserRepository interface {
		FindById(ctx context.Context, id int) (*database.User, error)
		FindByEmail(ctx context.Context, email string) (*database.User, error)
		Create(ctx context.Context, user *database.User) (int, error)
		Update(ctx context.Context) int
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

func (userRepo *userRepository) FindById(ctx context.Context, id int) (*database.User, error) {
	user, err := userRepo.db.FindById(ctx, 1)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepo *userRepository) FindByEmail(ctx context.Context, email string) (*database.User, error) {
	user, err := userRepo.db.FindByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepo *userRepository) Create(ctx context.Context, user *database.User) (int, error) {
	userCreate := database.CreateParams{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}
	_, err := userRepo.db.Create(ctx, userCreate)

	if err != nil {
		return 1, err
	}

	return 1, nil
}

func (userRepo *userRepository) Update(ctx context.Context) int {
	return 1
}
