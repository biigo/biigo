package biigo

import (
	"github.com/astaxie/beego/logs"
)

var logger *Logger

// Logger 扩展 BeeLogger
type Logger struct {
	*logs.BeeLogger
}

// Log 返回日志记录器
func Log() *Logger {
	if logger == nil {
		logger = &Logger{logs.NewLogger()}
	}
	return logger
}

// SetDefFileLog 设置默认的日志文件存储引擎
func (log *Logger) SetDefFileLog() *Logger {
	log.SetLogger(
		logs.AdapterMultiFile,
		`{
			"filename":"logs/biigo.log",
			"separate":[
				"emergency", "alert", "critical", "error", "warning"
			]
		}`,
	)
	return log
}
