package consts

type Verifycation int

const (
	TIME_OTP_REGISTER = 90 // 90 seconds

	EmailVerifycation Verifycation = 1
	PhoneVerifycation Verifycation = 2
)
