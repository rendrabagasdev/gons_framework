package mailer

import (
	"gons/internal/contracts"
	"gons/pkg/env"
	"strings"
)

// NewMailer is a factory function that returns the appropriate Mailer driver.
func NewMailer() contracts.Mailer {
	driver := env.Get("MAIL_MAILER", "log")

	switch strings.ToLower(driver) {
	case "smtp":
		return NewSMTPDriver()
	case "log":
		return NewLogDriver()
	default:
		// Default to log driver for safety
		return NewLogDriver()
	}
}
