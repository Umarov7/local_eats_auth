package service

import (
	"auth-service/config"
	"bytes"
	"fmt"
	"net/smtp"

	"github.com/pkg/errors"
)

func SendCode(cfg *config.Config, email string, code string) error {
	// sender data
	from := cfg.EMAIL
	password := cfg.APP_KEY

	// Receiver email address
	to := []string{email}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	var body bytes.Buffer

	body.Write([]byte(fmt.Sprintf("Subject: Your verification code \n%s\n\n", code)))

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return errors.Wrap(err, "failed to send email")
	}
	fmt.Println("Email sended to:", email)
	return nil
}
