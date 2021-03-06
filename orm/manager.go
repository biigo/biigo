package orm

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Manager 提供对数据库连接的管理
type Manager struct {
	config Config
	dbs    map[string]*gorm.DB
}

// Db 返回指定名称的数据库
func (manager *Manager) Db(name string) (*gorm.DB, error) {
	if manager.dbs == nil {
		manager.dbs = map[string]*gorm.DB{}
	}
	if db, ok := manager.dbs[name]; ok {
		return db, nil
	}
	config, ok := manager.config.Dbs[name]
	if !ok {
		return nil, fmt.Errorf("db %s not found", name)
	}
	db, err := gorm.Open(config.Driver(), config.URL)
	if err != nil {
		return nil, err
	}
	manager.dbs[name] = db
	return db, nil
}
