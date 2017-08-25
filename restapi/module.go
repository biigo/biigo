package restapi

import (
	"net/http"

	"github.com/biigo/biigo"
	"github.com/emicklei/go-restful"
	"github.com/go-openapi/spec"
)

// ModuleName 存储当前模块名称
const ModuleName = "restapi"

// NewModule 创建新的 restful 模块实例
func NewModule() *Module {
	return &Module{
		config:    Config{},
		ConfigKey: "restapi",
	}
}

// Module restful module
type Module struct {
	config    Config
	ConfigKey string
}

// ConfigApp 配置模块
func (module *Module) ConfigApp(app *biigo.App) error {
	return app.Config().JSONUnmarshal(module.ConfigKey, &module.config)
}

// InitApp 初始化应用程序
func (module *Module) InitApp(app *biigo.App) error {
	ws := new(restful.WebService)
	ws.Path(module.config.APIPrefix()).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	tags := []spec.Tag{}
	for _, module := range app.Modules() {
		if resContainer, ok := module.(ResourceContainer); ok {
			for _, resource := range resContainer.Resources() {
				resource.Register(ws)
			}
		}
		if tagContainer, ok := module.(ResourceTagContainer); ok {
			tags = append(tags, tagContainer.ResourceTags()...)
		}
	}

	restful.Add(ws)
	restful.Add(DocsResource{
		APIPath: module.config.Docs.DocsPath(),
		Tags:    tags,
		Info:    module.config.Docs.Info(),
	}.Register())
	return nil
}

// RunApp 运行接口服务器
func (module *Module) RunApp(errCh chan error) {
	listen := ":80"
	if module.config.HTTPListen != "" {
		listen = module.config.HTTPListen
	}
	err := http.ListenAndServe(listen, nil)
	if err != nil {
		errCh <- err
	}
}

// Name return module name
func (module *Module) Name() string {
	return ModuleName
}
