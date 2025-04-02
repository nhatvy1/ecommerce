package auth_strategies

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
	"user-service/pkg/response"
	"user-service/pkg/utils/auth"
	"user-service/pkg/utils/crypto"
	"user-service/pkg/utils/random"

	"github.com/redis/go-redis/v9"
)

type EmailRegisterStrategy struct {
	db *database.Queries
}

func NewEmailStrategy(db *database.Queries) *EmailRegisterStrategy {
	return &EmailRegisterStrategy{db: db}
}

func (e *EmailRegisterStrategy) Register(ctx context.Context, body *model.UserRegister) (int, int, error) {
	userFound, err := e.db.CheckUserBaseExists(ctx, body.UserAccount)
	if err != nil {
		return -1, http.StatusNotFound, err
	}

	if userFound > 0 {
		return -1, http.StatusConflict, fmt.Errorf("user has already registered")
	}

	hashAccount := crypto.GetHash(body.UserAccount)
	userKey := auth.SetKeyOTP(hashAccount)

	_, err = global.Rdb.Get(ctx, userKey).Result()

	switch {
	case err == redis.Nil:
		otpNew := random.GenerateOTP()
		err = global.Rdb.SetEx(ctx, userKey, strconv.Itoa(otpNew), time.Duration(consts.TIME_OTP_REGISTER)*time.Second).Err()
		if err != nil {
			return -1, http.StatusInternalServerError, fmt.Errorf("please enter again later")
		}
		fmt.Printf("----------")
		fmt.Println("User key: ", userKey)
		fmt.Println("Send otp: ", otpNew)
		fmt.Printf("----------")
		// err = send_otp.SendTextMailOtp([]string{body.UserAccount}, consts.HOST_EMAIL, strconv.Itoa(otpNew))
		// if err != nil {
		// 	return -1, http.StatusInternalServerError, fmt.Errorf("please enter again later")
		// }

		dataVerify := database.InsertUserVerifyParams{
			VerifyOtp:     strconv.Itoa(otpNew),
			VerifyKey:     body.UserAccount,
			VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
			VerifyKeyHash: hashAccount,
		}
		if err := e.db.InsertUserVerify(ctx, dataVerify); err != nil {
			fmt.Println(err)
			return -1, http.StatusBadRequest, fmt.Errorf("please try again later insert")
		}

		return 10, http.StatusOK, nil
	case err != nil:
		return response.ErrRedis, http.StatusInternalServerError, fmt.Errorf("please try again later")
	default:
		fmt.Println("User key: ", userKey)
		return response.OtpWaiting, http.StatusBadRequest, fmt.Errorf("please waiting 180 seconds")
	}
}
