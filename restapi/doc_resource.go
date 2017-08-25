package restapi

import (
	restful "github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
)

//DocsResource api docs resource
type DocsResource struct {
	Tags    []Tag
	APIPath string
	Info    *spec.Info
}

// SpecTags 返回接口标签
func (dr DocsResource) SpecTags() []spec.Tag {
	tags := []spec.Tag{}
	for _, tag := range dr.Tags {
		tags = append(tags, spec.Tag{TagProps: spec.TagProps{
			Name:        tag.Name,
			Description: tag.Desc,
		}})
	}
	return tags
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
	swo.Tags = dr.SpecTags()
}

// Tag 描述接口标签
type Tag struct {
	Name string
	Desc string
}
