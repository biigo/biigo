package restapi

import (
	restful "github.com/emicklei/go-restful"
)

// Resource 描述 restful 资源
type Resource interface {
	Register(*restful.WebService)
}

// ResourceContainer 描述资源容器
type ResourceContainer interface {
	Resources() []Resource
}

// ResourceTagContainer 描述资源标签容器
type ResourceTagContainer interface {
	ResourceTags() []Tag
}
