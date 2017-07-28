package biigo

// AppInitor 实用程序初始化器
type AppInitor interface {
	InitApp(app *App) error
}

type Module interface {
	Name() string
}

// App 描述应用程序
type App struct {
	modules map[string]Module
}

// Modules 返回已经注册的应用模块
func (app *App) Modules() map[string]Module {
	return app.modules
}

// AddModule 添加业务模块
func (app *App) AddModule(modules ...Module) *App {
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
