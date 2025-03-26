package utils

import (
	"context"
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(ctx context.Context, to, subject, body string) error {
	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")
	addr := os.Getenv("SMTP_ADDR")
	host := os.Getenv("SMTP_HOST")
	user := os.Getenv("EMAIL_USER")

	if from == "" || password == "" || addr == "" || host == "" {
		return fmt.Errorf("required environment variables are missing")
	}

	t := fmt.Sprintf("To: %s\r\n", to)
	s := fmt.Sprintf("Subject: %s\r\n", subject)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(t + s + mime + body + "\r\n")

	auth := smtp.PlainAuth("", user, password, host)

	done := make(chan error, 1)
	go func() {
		done <- smtp.SendMail(addr, auth, from, []string{to}, msg)
	}()

	select {
	case err := <-done:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return fmt.Errorf("email sending timed out")
	}

	return nil
}
