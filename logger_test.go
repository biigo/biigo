package biigo

import "testing"
import "github.com/astaxie/beego/logs"

func TestLogger(t *testing.T) {
	SetDefFileLog()
	Log().SetLogger(logs.AdapterConsole)
	Log().DelLogger(logs.AdapterMultiFile)
	Log().Trace("test log")
	Log().Error("test error log")
}
