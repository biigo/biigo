package mq

import (
	"github.com/biigo/biigo"
)

// ModuleName 存储当前模块名称
const ModuleName = "mq"

// ManagerSetter 定义需要依赖 MQ 的模块
type ManagerSetter interface {
	SetMqManager(*Manager)
}

// NewModule 创建一个 MQ 模块对象
func NewModule() *Module {
	return &Module{
		manager: &Manager{
			Config: Config{},
		},
	}
}

// Module mq 模块上下文
type Module struct {
	manager *Manager
}

// ConfigApp 配置模块
func (module *Module) ConfigApp(app *biigo.App) error {
	return app.Config().JSONUnmarshal("mq", &module.manager.Config)
}

// InitApp 初始化应用程序
func (module *Module) InitApp(app *biigo.App) error {
	for _, m := range app.Modules() {
		if mqSetter, ok := m.(ManagerSetter); ok {
			mqSetter.SetMqManager(module.manager)
		}
	}
	return nil
}

// Name return mq module name
func (module Module) Name() string {
	return ModuleName
}
