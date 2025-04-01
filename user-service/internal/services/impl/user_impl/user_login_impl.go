package user_impl

import (
	"context"
	"fmt"
	"net/http"
	"user-service/internal/database"
	"user-service/internal/model"
)

type userLogin struct {
	db              *database.Queries
	registerFactory *RegisterFactory
}

func NewUserLoginImpl(db *database.Queries, registerFactory *RegisterFactory) *userLogin {
	return &userLogin{
		db:              db,
		registerFactory: registerFactory,
	}
}

func (u *userLogin) Login(ctx context.Context) {}

func (u *userLogin) Register(ctx context.Context, body *model.UserRegister) (int, int, error) {
	// userFound, err := u.r.CheckUserBaseExists(ctx, body.UserAccount)
	// if err != nil {
	// 	return -1, http.StatusNotFound, err
	// }

	// if userFound > 0 {
	// 	return -1, http.StatusConflict, fmt.Errorf("user has already registered")
	// }

	// hashAccount := crypto.GetHash(body.UserAccount)
	// userKey := auth.SetKeyOTP(hashAccount)
	// otpNew := random.GenerateOTP()

	// err = global.Rdb.SetEx(ctx, userKey, strconv.Itoa(otpNew), time.Duration(consts.TIME_OTP_REGISTER)*time.Second).Err()
	// if err != nil {
	// 	return -1, http.StatusBadRequest, fmt.Errorf("please try again later")
	// }

	// dataVerify := database.InsertUserVerifyParams{
	// 	VerifyOtp:     strconv.Itoa(otpNew),
	// 	VerifyKey:     body.UserAccount,
	// 	VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
	// 	VerifyKeyHash: userKey,
	// }
	// if err := u.r.InsertUserVerify(ctx, dataVerify); err != nil {
	// 	return -1, http.StatusBadRequest, fmt.Errorf("please try again later")
	// }

	// 6. send OTP
	// fmt.Printf("User key: %s\n", userKey)
	// fmt.Printf("Send OTP: %s\n", strconv.Itoa(otpNew))
	strategy, err := u.registerFactory.GetStrategy(body)
	if err != nil {
		return -1, http.StatusBadRequest, fmt.Errorf("vui lòng thử lại sau")
	}

	return strategy.Execute(ctx, body)
}

func (u *userLogin) Verify(ctx context.Context) {}
