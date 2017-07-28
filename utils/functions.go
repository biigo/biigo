package utils

import (
	"errors"
	"log"
	"os"
	"path"
	"path/filepath"
)

//Getenv 获取环境变量
func Getenv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		v = def
	}
	return v
}

//FailOnError 显示错误消息后退出
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//ExtFiles 遍历根目录下所有指定后续的文件路径
func ExtFiles(root, ext string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(root, func(filePath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() || path.Ext(filePath) != ext {
			return nil
		}
		files = append(files, filePath)
		return nil
	})

	if err != nil {
		return []string{}, err
	}
	if len(files) < 1 {
		err = errors.New("没有找到任务符合条件的文件")
	}

	return files, err
}
