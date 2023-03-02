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

func InitAiderDateRouter(container *restful.Container) {
	s := &api.AiderDateApi{
		AiderDateApp: services.AiderDateModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/aider/date").Produces(restful.MIME_JSON)
	tags := []string{"aiderdate"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Date分页列表").Handle(s.GetAiderDateList)
	}).
		Doc("获取Date分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Date信息").Handle(s.GetAiderDate)
	}).
		Doc("获取Date信息").
		Param(ws.PathParameter("id", "Id").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.AiderDate{}). // on the response
		Returns(200, "OK", entity.AiderDate{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加Date信息").Handle(s.InsertAiderDate)
	}).
		Doc("添加Date信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.AiderDate{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改Date信息").Handle(s.UpdateAiderDate)
	}).
		Doc("修改Date信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.AiderDate{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除Date信息").Handle(s.DeleteAiderDate)
	}).
		Doc("删除Date信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
