package auth_strategies

import (
	"context"
	"fmt"
	"user-service/internal/database"
	"user-service/internal/model"
)

type EmailRegisterStrategy struct {
	db *database.Queries
}

func NewEmailStrategy(db *database.Queries) *EmailRegisterStrategy {
	return &EmailRegisterStrategy{db: db}
}

func (s *EmailRegisterStrategy) Register(ctx context.Context, body *model.UserRegister) (int, int, error) {
	fmt.Println("User register by email")
	return 1, 1, nil
}

// userFound, err := s.db.CheckUserBaseExists(ctx, body.UserAccount)
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
