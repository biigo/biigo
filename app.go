package biigo

// AppInitor 实用程序初始化器
type AppInitor interface {
	InitApp(app *App) error
}

// AppModule 描述业务模块
type AppModule interface {
	Name() string
}

// App 描述应用程序
type App struct {
	modules map[string]AppModule
	config  ConfigContainer
}

// Config 返回应用程序配置
func (app *App) Config() ConfigContainer {
	return app.config
}

// LoadConfig 从配置目录加载应用程序配置
func (app *App) LoadConfig(root string) *App {
	if config, err := ParseConfig(root); err != nil {
		panic(err)
	} else {
		app.config = config
	}
	return app
}

// Module 返回指定名称的模块
func (app *App) Module(name string) AppModule {
	m, ok := app.modules[name]
	if !ok {
		panic("Module " + name + " Not Found")
	}
	return m
}

// Modules 返回已经注册的应用模块
func (app *App) Modules() map[string]AppModule {
	return app.modules
}

// AddModule 添加业务模块
func (app *App) AddModule(modules ...AppModule) *App {
	for _, module := range modules {
		app.modules[module.Name()] = module
	}
	return app
}

// Init 初始化应用程序
func (app *App) Init() *App {
	for _, module := range app.modules {
		if appInitor, ok := module.(AppInitor); ok {
			if err := appInitor.InitApp(app); err != nil {
				panic(err)
			}
		}
	}
	return app
}
