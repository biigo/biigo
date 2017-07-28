package restful

import (
	"github.com/biigo/biigo"
	"github.com/emicklei/go-restful"
	"github.com/go-openapi/spec"
)

// Module restful module
type Module struct {
	APIPrefix  string
	APIDocPath string
	APIInfo    *spec.Info
}

// NewModule 创建新的 restful 模块实例
func NewModule(prefix, docsPath string, info *spec.Info) *Module {
	return &Module{
		APIPrefix:  prefix,
		APIDocPath: docsPath,
		APIInfo:    info,
	}
}

// InitApp 初始化应用程序
func (module *Module) InitApp(app *biigo.App) error {
	ws := new(restful.WebService)
	ws.Path(module.APIPrefix).
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

	if module.APIDocPath == "" {
		module.APIDocPath = "/api/docs.json"
	}

	restful.Add(ws)
	restful.Add(DocsResource{
		APIPath: module.APIDocPath,
		Tags:    tags,
		Info:    module.APIInfo,
	}.Register())
	return nil
}

// Name return module name
func (module *Module) Name() string {
	return "restfull"
}
