package biigo

var defApp = &App{
	modules: map[string]AppModule{},
}

// LoadConfig 加载应用程序配置
func LoadConfig(root string) *App {
	return defApp.LoadConfig(root)
}

// Config 返回系统配置
func Config() ConfigContainer {
	return defApp.Config()
}

// Module 返回指定名称的模块
func Module(name string) AppModule {
	return defApp.Module(name)
}

// Modules 返回已经注册的应用模块
func Modules() map[string]AppModule {
	return defApp.Modules()
}

// AddModule 添加业务模块
func AddModule(modules ...AppModule) *App {
	return defApp.AddModule(modules...)
}
