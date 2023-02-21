// ==========================================================================
// 生成日期：
// 生成路径: apps/aider/...
// 生成人：  ReWi
// ==========================================================================
package router

import (
	"GinMasterReWi/apps/aider/api"
	"GinMasterReWi/apps/aider/entity"
	"GinMasterReWi/apps/aider/services"
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitAiderProjectRouter(container *restful.Container) {
	s := &api.AiderProjectApi{
		AiderProjectApp: services.AiderProjectModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/aider/project").Produces(restful.MIME_JSON)
	tags := []string{"aiderproject"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Project分页列表").Handle(s.GetAiderProjectList)
	}).
		Doc("获取Project分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Project信息").Handle(s.GetAiderProject)
	}).
		Doc("获取Project信息").
		Param(ws.PathParameter("id", "Id").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.AiderProject{}). // on the response
		Returns(200, "OK", entity.AiderProject{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加Project信息").Handle(s.InsertAiderProject)
	}).
		Doc("添加Project信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.AiderProject{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改Project信息").Handle(s.UpdateAiderProject)
	}).
		Doc("修改Project信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.AiderProject{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除Project信息").Handle(s.DeleteAiderProject)
	}).
		Doc("删除Project信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	ws.Route(ws.POST("/upload").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("上传计划文件").Handle(s.UploadPlan)
	}).
		Doc("上传计划文件").
		Param(ws.FormParameter("planfile", "文件")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", map[string]string{}))

	container.Add(ws)
}
