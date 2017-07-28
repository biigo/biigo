package restful

import restful "github.com/emicklei/go-restful"

// Resource 描述 restful 资源
type Resource interface {
	Register(*restful.WebService)
}
