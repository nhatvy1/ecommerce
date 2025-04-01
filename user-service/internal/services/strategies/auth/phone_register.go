package auth_strategies

import (
	"context"
	"fmt"
	"user-service/internal/database"
	"user-service/internal/model"
)

type PhoneRegisterStrategy struct {
	db *database.Queries
}

func NewPhoneStrategy(db *database.Queries) *PhoneRegisterStrategy {
	return &PhoneRegisterStrategy{db: db}
}

func (s *PhoneRegisterStrategy) Register(ctx context.Context, body *model.UserRegister) (int, int, error) {
	fmt.Println("User register by phone")
	return 1, 1, nil
}
