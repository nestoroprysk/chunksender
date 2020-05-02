package sender

import (
	"sender/config"

	"gopkg.in/gomail.v2"
)

type EmailSender struct {
	dialer    *gomail.Dialer
	sender    string
	receivers []string
}

func NewEmailSender(ec config.EmailConfig) EmailSender {
	result := EmailSender{
		dialer: gomail.NewDialer(ec.Sender.Host, ec.Sender.Port,
			ec.Sender.Username, ec.Sender.Password),
		sender:    ec.Sender.Username,
		receivers: ec.Receivers,
	}

	return result
}

func (es EmailSender) Send(subject, msg string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", es.sender)
	m.SetHeader("To", es.receivers...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", msg)

	return es.dialer.DialAndSend(m)
}
