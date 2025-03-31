package auth

import (
	"fmt"
)

func SetKeyOTP(hash string) string {
	return fmt.Sprintf("u:%s:otp", hash)
}
