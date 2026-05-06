package mailer

import (
	"log/slog"
)

// LogDriver implements the Mailer interface by logging email content.
type LogDriver struct{}

// NewLogDriver creates a new instance of LogDriver.
func NewLogDriver() *LogDriver {
	return &LogDriver{}
}

// Send logs the email content instead of sending it.
func (l *LogDriver) Send(to string, subject string, body string) error {
	slog.Info("Email sent (Log Driver)",
		"to", to,
		"subject", subject,
		"body", body,
	)
	return nil
}
