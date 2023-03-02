package api

// ==========================================================================
// 生成日期：
// 生成路径: apps/aider/...
// 生成人：ReWi
// ==========================================================================
import (
	"GinMasterReWi/apps/aider/entity"
	"GinMasterReWi/apps/aider/services"
	"GinMasterReWi/pkg/global"
	"fmt"
	"github.com/XM-GO/PandaKit/biz"
	filek "github.com/XM-GO/PandaKit/file"
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"os"
	"time"

	"github.com/XM-GO/PandaKit/utils"

	"github.com/Re-Wi/GoKitReWi/handlers"
	"github.com/Re-Wi/GoKitReWi/helpers"
)

type AiderProjectApi struct {
	AiderProjectApp services.AiderProjectModel
}

// GetAiderProjectList 列表数据
func (p *AiderProjectApi) GetAiderProjectList(rc *restfulx.ReqCtx) {
	data := entity.AiderProject{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.Project = restfulx.QueryParam(rc, "project")

	list, total := p.AiderProjectApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetAiderProject 获取
func (p *AiderProjectApi) GetAiderProject(rc *restfulx.ReqCtx) {
	id := restfulx.PathParamInt(rc, "id")
	p.AiderProjectApp.FindOne(int64(id))
}

// InsertAiderProject 添加
func (p *AiderProjectApi) InsertAiderProject(rc *restfulx.ReqCtx) {
	var data entity.AiderProject
	restfulx.BindQuery(rc, &data)
	p.AiderProjectApp.Insert(data)
	rc.ResData = data // 返回值
}

// UpdateAiderProject 修改
func (p *AiderProjectApi) UpdateAiderProject(rc *restfulx.ReqCtx) {
	var data entity.AiderProject
	restfulx.BindQuery(rc, &data)
	p.AiderProjectApp.Update(data)
}

// DeleteAiderProject 删除
func (p *AiderProjectApi) DeleteAiderProject(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := utils.IdsStrToIdsIntGroup(id)
	p.AiderProjectApp.Delete(ids)
}

// UploadPlan 计划文件上传
func (p *AiderProjectApi) UploadPlan(rc *restfulx.ReqCtx) {
	_, fileHeader, err := rc.Request.Request.FormFile("planfile")
	fmt.Println("1", fileHeader.Filename)
	println("2", fileHeader)
	biz.ErrIsNil(err, "请传入文件")
	//Ext 函数返回 path 文件扩展名。返回值是路径最后一个斜杠分隔出的路径元素的最后一个 . 起始的后缀（包括 . ） 。如果该元素没有 . 会返回空字符串。
	//ext := path.Ext(fileHeader.Filename)
	//fmt.Println("3", ext)
	//TrimSuffix 返回没有提供的尾随后缀字符串的 s。如果 s 不以 suffix 结尾，则 s 原样返回。
	//// 读取文件名并加密
	//name := strings.TrimSuffix(fileHeader.Filename, ext)
	//fmt.Println("4", name)
	//name = kgo.KStr.Md5(name, 32)
	//fmt.Println("5", name)
	//// 拼接新文件名
	//filename := name + "_" + time.Now().Format("20060102150405") + ext
	//fmt.Println("6", filename)
	pwd, _ := os.Getwd()
	fmt.Println("7", pwd)
	filePath := fmt.Sprintf("%s/%s", helpers.CheckAndCreateDir(global.Conf.Upload.PlanDir+fmt.Sprintf("%d", time.Now().Year())), fileHeader.Filename)
	err = filek.SaveUploadedFile(fileHeader, filePath)
	fmt.Println("8", err)
	biz.ErrIsNil(err, "文件上传失败")

	var data entity.AiderProject
	filenameStruct := handlers.PlanFilenameParsing(fileHeader.Filename)
	data.Project = filenameStruct.Project
	data.StartTime = filenameStruct.StartTime
	data.FinishTime = filenameStruct.FinishTime
	data.Status = filenameStruct.Status
	plan, err := handlers.PlanContentParsing(filePath)
	data.Progress = 0
	data.LastTask = ""
	data.ThisTask = plan.PlanList[0].Plan

	p.AiderProjectApp.Insert(data)

	rc.ResData = map[string]string{"fileName": fileHeader.Filename}
}
