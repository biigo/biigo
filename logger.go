package biigo

import (
	"github.com/astaxie/beego/logs"
)

var logger *logs.BeeLogger

// Log 返回日志记录器
func Log() *logs.BeeLogger {
	if logger == nil {
		logger = logs.NewLogger()
	}
	return logger
}

// SetDefFileLog 设置默认的日志文件存储引擎
func SetDefFileLog() {
	Log().SetLogger(
		logs.AdapterMultiFile,
		`{
			"filename":"logs/biigo.log",
			"separate":[
				"emergency", "alert", "critical", "error", "warning"
			]
		}`,
	)
}
