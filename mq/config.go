package mq

// Config 定义模块配置
type Config struct {
	URLs map[string]string `json:"urls"`
}
