package services

import (
	"context"
	"user-service/internal/model"
)

type (
	IUserLogin interface {
		Login(ctx context.Context)
		Register(ctx context.Context, body *model.UserRegister) (int, int, error)
		Verify(ctx context.Context)
	}

	IUserInfo interface {
		GetInfoUserById(ctx context.Context)
	}

	IUserAdmin interface {
		GetListUsers(ctx context.Context)
	}

	IRegisterStrategy interface {
		Execute(ctx context.Context, body *model.UserRegister) (int, int, error)
	}
)

var (
	localUserLogin        IUserLogin
	localUserInfo         IUserInfo
	localUserAdmin        IUserAdmin
	localRegisterStrategy IRegisterStrategy
)

func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("implement user_login not found")
	}
	return localUserLogin
}

func InitUserLogin(i IUserLogin) {
	localUserLogin = i
}

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("implement user_info not found")
	}
	return localUserInfo
}

func InitUserInfo(i IUserInfo) {
	localUserInfo = i
}

func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("implement user_admin not found")
	}
	return localUserAdmin
}

func InitUserAdmin(i IUserAdmin) {
	localUserAdmin = i
}

func RegisterStrategy() IRegisterStrategy {
	if localRegisterStrategy == nil {
		panic("implement register_strategy not found")
	}
	return localRegisterStrategy
}

func InitRegisterStrategy(i IRegisterStrategy) {
	localRegisterStrategy = i
}
