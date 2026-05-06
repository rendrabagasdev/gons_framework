package mailer

import (
	"fmt"
	"gons/pkg/env"
	"net/smtp"
)

// SMTPDriver implements the Mailer interface using net/smtp.
type SMTPDriver struct {
	host     string
	port     string
	username string
	password string
	from     string
}

// NewSMTPDriver creates a new instance of SMTPDriver.
func NewSMTPDriver() *SMTPDriver {
	return &SMTPDriver{
		host:     env.Get("MAIL_HOST", "localhost"),
		port:     env.Get("MAIL_PORT", "1025"),
		username: env.Get("MAIL_USERNAME", ""),
		password: env.Get("MAIL_PASSWORD", ""),
		from:     env.Get("MAIL_FROM_ADDRESS", "hello@example.com"),
	}
}

// Send sends an email using SMTP.
func (s *SMTPDriver) Send(to string, subject string, body string) error {
	auth := smtp.PlainAuth("", s.username, s.password, s.host)
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", to, subject, body))
	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	return smtp.SendMail(addr, auth, s.from, []string{to}, msg)
}
