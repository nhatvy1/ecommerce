package user_impl

import (
	"context"
	"database/sql"
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
	hashAccount := crypto.GetHash(body.UserAccount)

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
	if err != nil {
		return -1, http.StatusBadRequest, fmt.Errorf("please try again later")
	}

	dataVerify := database.InsertUserVerifyParams{
		VerifyOtp:     strconv.Itoa(otpNew),
		VerifyKey:     body.UserAccount,
		VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
		VerifyKeyHash: userKey,
	}
	if err := u.r.InsertUserVerify(ctx, dataVerify); err != nil {
		return -1, http.StatusBadRequest, fmt.Errorf("please try again later")
	}

	// 6. send OTP
	fmt.Printf("User key: %s\n", userKey)
	fmt.Printf("Send OTP: %s\n", strconv.Itoa(otpNew))

	// return 1, http.StatusBadRequest, fmt.Errorf("not found")
	return 1, http.StatusOK, nil
}

func (u *userLogin) Verify(ctx context.Context) {}
