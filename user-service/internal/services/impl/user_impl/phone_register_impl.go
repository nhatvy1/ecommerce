package user_impl

import (
	"context"
	"fmt"
	"user-service/internal/database"
	"user-service/internal/model"
)

type PhoneStrategy struct {
	db *database.Queries
}

func NewPhoneStrategy(db *database.Queries) *PhoneStrategy {
	return &PhoneStrategy{db: db}
}

func (s *PhoneStrategy) Execute(ctx context.Context, body *model.UserRegister) (int, int, error) {
	fmt.Println("User register by phone")
	return 1, 1, nil
}
