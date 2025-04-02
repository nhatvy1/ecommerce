package send_otp

import (
	"fmt"
	"net/smtp"
	"strings"
	"user-service/global"

	"go.uber.org/zap"
)

const (
	SMTPHost     = "smtp.gmail.com"
	SMTPPort     = "587"
	SMTPFrom     = "vyhuynhak123@gmail.com"
	SMTPPassword = "lrndszovebyhpgkd"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func SendTextMailOtp(to []string, from string, otp string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "test"},
		To:      to,
		Subject: "OTP Verification",
		Body:    fmt.Sprintf("Your OTP is %s. Please enter it to verify your account.", otp),
	}

	messageMail := BuildMessage(contentEmail)

	//send smtp
	auth := smtp.PlainAuth("", SMTPFrom, SMTPPassword, SMTPHost)

	err := smtp.SendMail(SMTPHost+":25", auth, from, to, []byte(messageMail))
	if err != nil {
		global.Logger.Error("Email send failed::", zap.Error(err))
		return err
	}

	return nil
}
