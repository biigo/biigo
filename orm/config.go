package orm

// Config 定义模块配置
type Config struct {
	Dbs []DbConfig `json:"dbs"`
}

// DbConfig 定义数据库连接配置
type DbConfig struct {
	DriverName string `json:"driver"`
	URL        string `json:"url"`
	ConfName   string `json:"name"`
}

// Driver return db driver name
func (config DbConfig) Driver() string {
	if config.DriverName == "" {
		return "mysql"
	}
	return config.DriverName
}

// Name return db config name
func (config DbConfig) Name() string {
	if config.ConfName == "" {
		return "default"
	}
	return config.ConfName
}
