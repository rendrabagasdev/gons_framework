package contracts

// Mailer defines the contract for sending emails.
type Mailer interface {
	Send(to string, subject string, body string) error
}
