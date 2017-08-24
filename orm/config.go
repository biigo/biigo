package orm

import "errors"

// Config 定义模块配置
type Config struct {
	Dbs map[string]DbConfig `json:"dbs"`
}

// Valid 验证配置是否正确
func (c Config) Valid() error {
	if len(c.Dbs) < 1 {
		return errors.New("orm config error")
	}
	return nil
}

// DbConfig 定义数据库连接配置
type DbConfig struct {
	DriverName string `json:"driver"`
	URL        string `json:"url"`
}

// Driver return db driver name
func (config DbConfig) Driver() string {
	if config.DriverName == "" {
		return "mysql"
	}
	return config.DriverName
}
