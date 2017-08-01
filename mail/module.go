package mail

// ModuleName 存储当前模块名称
const ModuleName = "mail"

var context = &Context{}

// Context mail 模块上下文
type Context struct {
	defSender Sender
}

// Module 返回当前模块实例
func Module(defSender Sender) *Context {
	context.defSender = defSender
	return context
}

// Sender 返回默认邮件发送器
func (context *Context) Sender() Sender {
	return context.defSender
}

// Name return mail module name
func (context *Context) Name() string {
	return ModuleName
}
