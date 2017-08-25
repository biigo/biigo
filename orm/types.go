package orm

import (
	"time"
)

// ManagerSetter 定义需要依赖数据库的模块
type ManagerSetter interface {
	SetOrmManager(*Manager) error
}

// DataMigrator 定义数据迁移器
type DataMigrator interface {

	// AutoMigrate 运行自动迁移
	AutoMigrate() error
}

//Model base model
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"deletedAt"`
}
