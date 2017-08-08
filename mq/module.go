package mq

import (
	"errors"

	"github.com/biigo/biigo"
	"github.com/streadway/amqp"
)

// ModuleName 存储当前模块名称
const ModuleName = "mq"

var context = &Context{
	configs: []config{},
}

// MqConnectionSetter 定义需要依赖 MQ 连接的模块
type MqConnectionSetter interface {
	SetMQConnect(*amqp.Connection)
}

// Context mq 模块上下文
type Context struct {
	configs []config
	defConn *amqp.Connection
}

// Module return context
func Module() *Context {
	return context
}

// ConfigApp 配置模块
func (context *Context) ConfigApp(app *biigo.App) error {
	return app.Config().JSONUnmarshal("mq", &context.configs)
}

// InitApp 初始化应用程序
func (context *Context) InitApp(app *biigo.App) error {
	conn, err := context.DefConn()
	if err != nil {
		return err
	}
	for _, module := range app.Modules() {
		if connSetter, ok := module.(MqConnectionSetter); ok {
			connSetter.SetMQConnect(conn)
		}
	}
	return nil
}

// Name return mq module name
func (context *Context) Name() string {
	return ModuleName
}

// DefConn 返回默认消息中间件连接
func (context *Context) DefConn() (*amqp.Connection, error) {
	if context.defConn == nil && len(context.configs) < 1 {
		return context.defConn, errors.New("缺少消息中间件连接配置")
	}
	if context.defConn == nil {
		if conn, err := amqp.Dial(context.configs[0].URI); err == nil {
			context.defConn = conn
		} else {
			return context.defConn, err
		}
	}
	return context.defConn, nil
}

type config struct {
	URI string `json:"uri"`
}
