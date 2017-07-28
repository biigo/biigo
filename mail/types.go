package mail

// Message 描述邮件内容信息
type Message struct {
	From     string
	To       []string
	Cc       []string
	Subject  string
	BodyMime string
	Body     string
}

// Sender 描述邮件发送器
type Sender interface {
	Send(Message) error
}
