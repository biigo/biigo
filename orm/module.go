package orm

import (
	"github.com/biigo/biigo"
	"github.com/jinzhu/gorm"
)

// NewModule 创建模块实例
func NewModule() *Module {
	module := &Module{
		ConfigKeyName: "orm",
		Manager: &Manager{
			config: Config{},
			dbs:    map[string]*gorm.DB{},
		},
	}
	return module
}

// Module orm module
type Module struct {
	ConfigKeyName string
	Manager       *Manager
}

// ConfigApp 配置模块
func (module *Module) ConfigApp(app *biigo.App) error {
	return app.Config().JSONUnmarshal(
		module.ConfigKeyName,
		&module.Manager.config)
}

// InitApp 初始化应用程序
func (module *Module) InitApp(app *biigo.App) error {
	for _, m := range app.Modules() {
		if managerSetter, ok := m.(ManagerSetter); ok {
			if err := managerSetter.SetOrmManager(module.Manager); err != nil {
				return err
			}
		}
	}
	return nil
}

// Name return orm module name
func (module *Module) Name() string {
	return "orm"
}
