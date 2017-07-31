package restapi

import (
	"log"
	"net/http"

	"github.com/biigo/biigo"
	"github.com/emicklei/go-restful"
	"github.com/go-openapi/spec"
)

// ModuleName 存储当前模块名称
const ModuleName = "restapi"

// Module restful module
type Module struct {
	APIPrefix  string
	APIDocPath string
	APIInfo    *spec.Info
	HTTPListen string
}

// NewModule 创建新的 restful 模块实例
func NewModule(prefix, docsPath string, info *spec.Info, httpListen ...string) *Module {
	module := &Module{
		APIPrefix:  prefix,
		APIDocPath: docsPath,
		APIInfo:    info,
	}

	if len(httpListen) > 0 {
		module.HTTPListen = httpListen[0]
	}

	return module
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

// RunWebServer 运行接口服务器
func (module *Module) RunWebServer() {
	if module.HTTPListen == "" {
		module.HTTPListen = ":80"
	}
	log.Fatal(http.ListenAndServe(module.HTTPListen, nil))
}

// Name return module name
func (module *Module) Name() string {
	return ModuleName
}
