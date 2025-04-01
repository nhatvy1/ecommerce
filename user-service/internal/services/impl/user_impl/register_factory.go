package user_impl

import (
	"errors"
	"user-service/internal/database"
	"user-service/internal/model"
	"user-service/internal/services"
)

type RegisterFactory struct {
	strategies map[string]services.IRegisterStrategy
}

func NewRegisterFactory(db *database.Queries) *RegisterFactory {
	return &RegisterFactory{
		strategies: map[string]services.IRegisterStrategy{
			"email": NewEmailStrategy(db),
			"phone": NewPhoneStrategy(db),
		},
	}
}

func (f *RegisterFactory) GetStrategy(body *model.UserRegister) (services.IRegisterStrategy, error) {
	switch {
	case body.AccountType != "1":
		return f.strategies["email"], nil
	case body.AccountType != "2":
		return f.strategies["phone"], nil
	default:
		return nil, errors.New("unsupported registration method")
	}
}
