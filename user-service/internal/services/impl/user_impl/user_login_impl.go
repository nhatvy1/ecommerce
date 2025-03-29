package user_impl

import (
	"context"
	"fmt"
	"net/http"
	"user-service/internal/database"
	"user-service/internal/model"
)

type userLogin struct {
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *userLogin {
	return &userLogin{
		r: r,
	}
}

func (u *userLogin) Login(ctx context.Context) {}

func (u *userLogin) Register(ctx context.Context, body *model.UserRegister) (int, int, error) {
	// 1. hash email
	// 2. check user exists
	// 3. create OTP
	// 4. generate OTP
	// 5. save OTP in redis with expiration time
	// 6. send OTP

	return 1, http.StatusBadRequest, fmt.Errorf("not found")
}

func (u *userLogin) Verify(ctx context.Context) {}
