package mail

import (
	"github.com/stretchr/testify/mock"
)

// MockMailSender 模拟邮件发送器
type MockMailSender struct {
	mock.Mock
}

// Send 模拟发送邮件
func (mock *MockMailSender) Send(msg Message) error {
	args := mock.Called(msg)
	return args.Error(0)
}
