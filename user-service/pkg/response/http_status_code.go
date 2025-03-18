package response

const (
	ErrCodeSuccess       = 20001 // Success
	ErrCodeParamInvalid  = 20003 // Email is invalid
	ErrInvalidToken      = 30001 // Token is invalid
	ErrPlsTryAgainLater  = 40003
	ErrCodeUserHasExists = 50001 // User has already registered,
	ErrCodeAuthFailed    = 40005
)

var Msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "email is invalid",
	ErrInvalidToken:      "token is invalid",
	ErrPlsTryAgainLater:  "please try again later",
	ErrCodeUserHasExists: "user has already registered",
	ErrCodeAuthFailed:    "Authentication failed",
}
