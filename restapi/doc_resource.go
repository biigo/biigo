package restapi

import (
	restful "github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
)

//DocsResource api docs resource
type DocsResource struct {
	Tags    []spec.Tag
	APIPath string
	Info    *spec.Info
}

//Register api docs resource
func (dr DocsResource) Register() *restful.WebService {
	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(),
		APIPath:     dr.APIPath,
		PostBuildSwaggerObjectHandler: dr.enrichSwaggerObject,
	}
	return restfulspec.NewOpenAPIService(config)
}

func (dr DocsResource) enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{}
	swo.Tags = dr.Tags
}
