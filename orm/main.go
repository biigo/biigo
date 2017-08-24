package orm

import (
	"github.com/jinzhu/gorm"
)

var manager = DbManager{
	dbs: map[string]*gorm.DB{},
}

// InitOrm 初始化 ORM 服务
func InitOrm(config Config) error {
	if err := config.Valid(); err != nil {
		return err
	}
	manager.config = config
	return nil
}

// Db 返回定名称的数据库
func Db(name string) (*gorm.DB, error) {
	return manager.Db(name)
}
