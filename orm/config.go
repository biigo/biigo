package orm

// Config 定义模块配置
type Config struct {
	Dbs map[string]DbConfig `json:"dbs"`
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
