package srv

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

type EmailService struct {
	auth smtp.Auth
}

const (
	EmailAddrFrom = "storichallenge1@gmail.com"
)

func CreateEmailService() (EmailService, error) {
	password := strings.TrimSpace(os.Getenv("EMAIL_PASSWORD"))
	if password == "" {
		return EmailService{}, fmt.Errorf("EMAIL_PASSWORD environment variable not set")
	}

	auth := smtp.PlainAuth("", EmailAddrFrom, password, "smtp.gmail.com")

	return EmailService{
		auth: auth,
	}, nil
}

func (s *EmailService) SendEmail(emailAddrTo, body string) error {
	to := emailAddrTo
	subject := "Transactions Balance"
	message := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	err := smtp.SendMail("smtp.gmail.com:587", s.auth, EmailAddrFrom, []string{to}, message)
	if err != nil {
		return err
	}

	return nil
}
