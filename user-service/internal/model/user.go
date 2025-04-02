package model

type UserRegister struct {
	UserAccount string `json:"user_account" validate:"required,email"`
	AccountType string `json:"account_type"`
}

type VerifyInput struct {
	UserAccount string `json:"user_account"`
	VerifyCode  string `json:"verify_code"`
}

type VerifyOtpOutput struct {
	Token   string `json:"token"`
	UserId  string `json:"user_id"`
	Message string `json:"message"`
}
