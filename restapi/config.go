package restapi

import "github.com/go-openapi/spec"

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

// DocsPath 返回接口文档地址
func (c DocConfig) DocsPath() string {
	if c.Path == "" {
		c.Path = "/api/docs.json"
	}
	return c.Path
}

// Info 返回接口文档信息
func (c DocConfig) Info() *spec.Info {
	return &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       c.Title,
			Description: c.Desc,
			Contact: &spec.ContactInfo{
				Name:  c.ContactName,
				Email: c.ContactEmail,
			},
			License: &spec.License{},
			Version: c.Version,
		},
	}
}
