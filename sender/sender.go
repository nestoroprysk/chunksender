package sender

import "github.com/nestoroprysk/chunksender/config"

type Sender interface {
	Send(subject, msg string) error
}

func NewSenders(cfg config.Config) []Sender {
	result := []Sender{}
	if cfg.Email != nil {
		var emailSender Sender = NewEmailSender(*cfg.Email)
		result = append(result, emailSender)
	}

	return result
}
