package user_impl

import (
	"context"
	"fmt"
	"user-service/global"
	"user-service/internal/database"
	"user-service/internal/model"
	"user-service/internal/services"
	auth_strategies "user-service/internal/services/strategies/auth"
	"user-service/pkg/utils/auth"
	"user-service/pkg/utils/crypto"
)

type userLogin struct {
	db                 *database.Queries
	registerStrategies map[string]services.IAuthStrategy
}

func NewUserLoginImpl(db *database.Queries) *userLogin {
	return &userLogin{
		db: db,
		registerStrategies: map[string]services.IAuthStrategy{
			"email": auth_strategies.NewEmailStrategy(db),
			"phone": auth_strategies.NewPhoneStrategy(db),
		},
	}
}

func (u *userLogin) Login(ctx context.Context) {}

func (u *userLogin) Register(ctx context.Context, body *model.UserRegister) (int, int, error) {
	accountType := body.AccountType
	var accountTypename string

	switch accountType {
	case "1":
		accountTypename = "email"
	case "2":
		accountTypename = "phone"
	default:
		accountTypename = "unknown"
	}

	strategy, ok := u.registerStrategies[accountTypename]
	if !ok {
		strategy = &auth_strategies.DefaultRegisterStrategy{}
	}

	return strategy.Register(ctx, body)
}

func (u *userLogin) VerifyOTP(ctx context.Context, body *model.VerifyInput) (out model.VerifyOtpOutput, err error) {
	hashAccount := crypto.GetHash(body.UserAccount)
	userKey := auth.SetKeyOTP(hashAccount)

	otpFound, err := global.Rdb.Get(ctx, userKey).Result()
	if err != nil {
		return out, err
	}

	if body.VerifyCode != otpFound {
		return out, fmt.Errorf("OTP not match")
	}

	out.Token = infoOtp.VerifyKeyHash
	out.Message = "success"

	return out, nil
}
