package mailer

type Client interface {
	Send(templateFile string, username string, email string, data any, isSandbox bool) error
}
