package model

type LoginOutput struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
