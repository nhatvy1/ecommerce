package user_impl

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"user-service/global"
	consts "user-service/internal/const"
	"user-service/internal/database"
	"user-service/internal/model"
	"user-service/pkg/utils/auth"
	"user-service/pkg/utils/crypto"
	"user-service/pkg/utils/random"
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
	hashAccount := crypto.GetHash(body.UserAccount)

	// 2. check user exists
	userFound, err := u.r.CheckUserBaseExists(ctx, body.UserAccount)
	if err != nil {
		return -1, http.StatusNotFound, err
	}

	if userFound > 0 {
		return -1, http.StatusConflict, fmt.Errorf("user has already registered")
	}

	userKey := auth.SetKeyOTP(hashAccount)
	otpNew := random.GenerateOTP()
	err = global.Rdb.SetEx(ctx, userKey, strconv.Itoa(otpNew), time.Duration(consts.TIME_OTP_REGISTER)*time.Second).Err()

	fmt.Println(err)
	fmt.Println("user key: ", userKey)
	return 1, http.StatusOK, nil

	// 3. create OTP
	// 4. generate OTP
	// 5. save OTP in redis with expiration time
	// 6. send OTP

	// return 1, http.StatusBadRequest, fmt.Errorf("not found")
}

func (u *userLogin) Verify(ctx context.Context) {}
