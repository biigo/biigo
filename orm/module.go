package orm

import (
	"github.com/biigo/biigo"
	"github.com/jinzhu/gorm"
)

// Module orm module
type Module struct {
	dirver string
	url    string
	DB     *gorm.DB
}

// NewModule 创建模块实例
func NewModule(dirver, url string) *Module {
	return &Module{
		dirver: dirver,
		url:    url,
	}
}

// InitApp 初始化应用程序
func (module *Module) InitApp(app *biigo.App) error {
	db, err := gorm.Open(module.dirver, module.url)
	if err != nil {
		return err
	}
	module.DB = db

	tx := db.Begin()
	for _, module := range app.Modules() {
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
