package orm

import (
	"time"

	"github.com/jinzhu/gorm"
)

// DataMigrator 定义数据迁移器
type DataMigrator interface {

	// AutoMigration 运行自动迁移
	AutoMigration(*gorm.DB) error
}

//Model base model
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"deletedAt"`
}
