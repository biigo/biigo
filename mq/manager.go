package mq

import (
	"errors"

	"github.com/streadway/amqp"
)

// Manager 定义MQ连接管理器
type Manager struct {
	Config Config
	conns  map[string]*amqp.Connection
}

// Conn 返回指定名称的连接对象
func (m *Manager) Conn(name string) (*amqp.Connection, error) {
	if conn, ok := m.conns[name]; ok {
		return conn, nil
	}
	url, ok := m.Config.URLs[name]
	if !ok {
		return nil, errors.New("mq connection not found")
	}
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	m.conns[name] = conn
	return m.conns[name], nil
}
