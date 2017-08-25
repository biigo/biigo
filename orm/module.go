package orm

import (
	"fmt"

	"github.com/biigo/biigo"
	"github.com/jinzhu/gorm"
)

// NewModule 创建模块实例
func NewModule() *Module {
	module := &Module{
		config:        Config{},
		ConfigKeyName: "orm",
		dbs:           map[string]*gorm.DB{},
	}
	return module
}

// Module orm module
type Module struct {
	config        Config
	ConfigKeyName string
	dbs           map[string]*gorm.DB
}

// ConfigApp 配置模块
func (module *Module) ConfigApp(app *biigo.App) error {
	return app.Config().JSONUnmarshal(module.ConfigKeyName, &module.config)
}

// Db 返回指定名称的数据库
func (module *Module) Db(name string) (*gorm.DB, error) {
	if db, ok := module.dbs[name]; ok {
		return db, nil
	}
	config, ok := module.config.Dbs[name]
	if !ok {
		return nil, fmt.Errorf("db %s not found", name)
	}
	db, err := gorm.Open(config.Driver(), config.URL)
	if err != nil {
		return nil, err
	}
	module.dbs[name] = db
	return db, nil
}

// InitApp 初始化应用程序
func (module *Module) InitApp(app *biigo.App) error {
	db, err := module.Db("default")
	if err != nil {
		return err
	}

	tx := db.Begin()
	for _, module := range app.Modules() {
		if dbWriter, ok := module.(DbWriter); ok {
			dbWriter.SetDB(db)
		}
		if dataMigrator, ok := module.(DataMigrator); ok {
			if err := dataMigrator.AutoMigrate(tx); err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	tx.Commit()
	return nil
}

// Name return orm module name
func (module *Module) Name() string {
	return "orm"
}
