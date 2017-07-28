package mail

import (
	"errors"

	gomail "gopkg.in/gomail.v2"
)

// GomailSender gomail 的 sender 实现
type GomailSender struct {
	Dialer  *gomail.Dialer
	DefFrom string
	DefTo   []string
}

// Send 发送邮件
func (sender GomailSender) Send(msg Message) error {
	from := sender.DefFrom
	to := sender.DefTo
	if msg.From != "" {
		from = msg.From
	}
	if len(msg.To) > 0 {
		to = msg.To
	}
	if from == "" || len(to) < 1 {
		return errors.New("邮件发送者或接收者不能为空")
	}
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", msg.Subject)
	m.SetBody(msg.BodyMime, msg.Body)
	if len(msg.Cc) > 0 {
		m.SetHeader("Cc", msg.Cc...)
	}

	return sender.Dialer.DialAndSend(m)
}
