package restapi

// Config 定义restapi模块配置
type Config struct {
	Prefix     string    `json:"apiPrefix"`
	HTTPListen string    `jsn:"apiHttpListen"`
	Docs       DocConfig `json:"apiDocs"`
}

// DocConfig 描述接口文档配置
type DocConfig struct {
	Path         string `json:"path"`
	Title        string `json:"title"`
	Desc         string `json:"desc"`
	ContactName  string `json:"contactName"`
	ContactEmail string `json:"contactEmail"`
	Version      string `json:"version"`
}

// APIPrefix 返回回接口前缀
func (c Config) APIPrefix() string {
	if c.Prefix == "" {
		return "/api"
	}
	return c.Prefix
}
