package middleware

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/go-openapi/spec"
)

/**
 * @Description
 * @Author 熊猫
 * @Date 2022/8/3 9:16
 **/

func SwaggerConfig(wsContainer *restful.Container) {
	config := restfulspec.Config{
		WebServices:                   wsContainer.RegisteredWebServices(),
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	wsContainer.Add(restfulspec.NewOpenAPIService(config))
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "GinMasterReWi 框架的API文档",
			Description: "这是GinMasterReWi框架文档，根据文档调用API",
			Contact: &spec.ContactInfo{
				ContactInfoProps: spec.ContactInfoProps{
					Name:  "GinMasterReWi 总怿窗",
					Email: "1871002801@qq.com",
					URL:   "https://github.com/Re-Wi/MasterReWi",
				},
			},
			License: &spec.License{
				LicenseProps: spec.LicenseProps{
					Name: "MIT",
					URL:  "https://github.com/Re-Wi/MasterReWi",
				},
			},
			Version: "1.0.0",
		},
	}
}
