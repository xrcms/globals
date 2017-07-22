package globals

type Emailer interface {
	Send(recipientEmail, recipientName, subject, message string, files ...string) error
}
