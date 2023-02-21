package router

import (
	"GinMasterReWi/apps/system/api"
	"github.com/emicklei/go-restful/v3"
)

func InitSystemRouter(container *restful.Container) {
	s := &api.System{}
	ws := new(restful.WebService)
	ws.Path("/system").Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(s.ConnectWs))
	ws.Route(ws.GET("/server").To(s.ServerInfo))
	container.Add(ws)
}
