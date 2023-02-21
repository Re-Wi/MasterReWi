package api

// ==========================================================================
// 生成日期：
// 生成路径: apps/aider/...
// 生成人：ReWi
// ==========================================================================
import (
	"GinMasterReWi/apps/aider/entity"
	"GinMasterReWi/apps/aider/services"
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/XM-GO/PandaKit/utils"
)

type AiderDateApi struct {
	AiderDateApp services.AiderDateModel
}

// GetAiderDateList 列表数据
func (p *AiderDateApi) GetAiderDateList(rc *restfulx.ReqCtx) {
	data := entity.AiderDate{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)

	list, total := p.AiderDateApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetAiderDate 获取
func (p *AiderDateApi) GetAiderDate(rc *restfulx.ReqCtx) {
	id := restfulx.PathParamInt(rc, "id")
	p.AiderDateApp.FindOne(int64(id))
}

// InsertAiderDate 添加
func (p *AiderDateApi) InsertAiderDate(rc *restfulx.ReqCtx) {
	var data entity.AiderDate
	restfulx.BindQuery(rc, &data)

	p.AiderDateApp.Insert(data)
}

// UpdateAiderDate 修改
func (p *AiderDateApi) UpdateAiderDate(rc *restfulx.ReqCtx) {
	var data entity.AiderDate
	restfulx.BindQuery(rc, &data)

	p.AiderDateApp.Update(data)
}

// DeleteAiderDate 删除
func (p *AiderDateApi) DeleteAiderDate(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := utils.IdsStrToIdsIntGroup(id)
	p.AiderDateApp.Delete(ids)
}
