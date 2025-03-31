package model

type UserRegister struct {
	UserAccount string `json:"user_account" validate:"required,email"`
	AccountType string `json:"account_type"`
}
