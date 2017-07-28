package biigo

//Config service
type Config struct {
	Def map[string]interface{}
}

//Get config
func (c Config) Get(key string, def interface{}) interface{} {
	if v, ok := c.Def[key]; ok {
		return v
	}
	return def
}

// String 返回字符串配置
func (c Config) String(key, def string) string {
	return c.Get(key, def).(string)
}
