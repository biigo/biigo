package biigo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/biigo/biigo/utils"
)

// ParseConfig 加载配置
func ParseConfig(root string) (Config, error) {
	config := Config{Values: make(map[string]interface{})}

	files, err := utils.ExtFiles(root, ".json")
	if err != nil {
		return config, errors.New("加载配置文件列表失败：" + err.Error())
	}

	localFiles := []string{}
	for _, f := range files {
		if isLocal, _ := path.Match(".*\\.local\\.json$", f); isLocal {
			localFiles = append(localFiles, f)
			continue
		}
		if content, err := ioutil.ReadFile(f); err != nil {
			return config, fmt.Errorf("加载配置文件失败：%s %s", f, err.Error())
		} else {
			conf := map[string]interface{}{}
			if err := json.Unmarshal(content, &conf); err != nil {
				return config, fmt.Errorf("解析配置文件失败：%s %s", f, err.Error())
			}
			for name, value := range conf {
				config.Values[name] = value
			}
		}
	}
	if len(localFiles) > 0 {
		for _, f := range localFiles {
			if content, err := ioutil.ReadFile(f); err != nil {
				return config, fmt.Errorf("加载配置文件失败：%s %s", f, err.Error())
			} else {
				conf := map[string]interface{}{}
				if err := json.Unmarshal(content, &conf); err != nil {
					return config, fmt.Errorf("解析配置文件失败：%s %s", f, err.Error())
				}
				for name, value := range conf {
					config.Values[name] = value
				}
			}
		}
	}
	return config, nil
}

//Config service
type Config struct {
	Values map[string]interface{}
}

//Get config
func (c Config) Get(key string, def interface{}) interface{} {
	if v, ok := c.Values[key]; ok {
		return v
	}
	return def
}

// String 返回字符串配置
func (c Config) String(key, def string) string {
	return c.Get(key, def).(string)
}

// JSONUnmarshal 利用 json.Unmarshal 解析配置
func (c Config) JSONUnmarshal(key string, v interface{}) error {
	content := c.Get(key, nil)
	if content == nil {
		return fmt.Errorf("config key %s not found", key)
	}
	bytes, err := json.Marshal(content)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, v)
}
