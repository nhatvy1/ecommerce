package user_impl

import (
	"context"
	"fmt"
	"user-service/internal/database"
	"user-service/internal/model"
)

type EmailStrategy struct {
	db *database.Queries
}

func NewEmailStrategy(db *database.Queries) *EmailStrategy {
	return &EmailStrategy{db: db}
}

func (s *EmailStrategy) Execute(ctx context.Context, body *model.UserRegister) (int, int, error) {
	fmt.Println("User register by email")
	return 1, 1, nil
}
