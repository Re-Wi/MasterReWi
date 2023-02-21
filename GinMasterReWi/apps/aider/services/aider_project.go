// ==========================================================================
// 生成日期：
// 生成路径: apps/aider/...
// 生成人：ReWi
// ==========================================================================

package services

import (
	"GinMasterReWi/apps/aider/entity"
	"GinMasterReWi/pkg/global"
	"github.com/XM-GO/PandaKit/biz"
)

type (
	AiderProjectModel interface {
		Insert(data entity.AiderProject) *entity.AiderProject
		FindOne(id int64) *entity.AiderProject
		FindListPage(page, pageSize int, data entity.AiderProject) (*[]entity.AiderProject, int64)
		FindList(data entity.AiderProject) *[]entity.AiderProject
		Update(data entity.AiderProject) *entity.AiderProject
		Delete(ids []int64)
	}

	projectModelImpl struct {
		table string
	}
)

var AiderProjectModelDao AiderProjectModel = &projectModelImpl{
	table: `aider_projects`,
}

func (m *projectModelImpl) Insert(data entity.AiderProject) *entity.AiderProject {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加项目失败")
	return &data
}

func (m *projectModelImpl) FindOne(id int64) *entity.AiderProject {
	resData := new(entity.AiderProject)
	err := global.Db.Table(m.table).Where("id = ?", id).First(resData).Error
	biz.ErrIsNil(err, "查询项目失败")
	return resData
}

func (m *projectModelImpl) FindListPage(page, pageSize int, data entity.AiderProject) (*[]entity.AiderProject, int64) {
	list := make([]entity.AiderProject, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Project != "" {
		db = db.Where("name like ?", "%"+data.Project+"%")
	}
	//if data.Creator != 0 {
	//	db = db.Where("creator = ?", data.Creator)
	//}
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询项目分页列表失败")
	return &list, total
}

func (m *projectModelImpl) FindList(data entity.AiderProject) *[]entity.AiderProject {
	list := make([]entity.AiderProject, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Project != "" {
		db = db.Where("name like ?", "%"+data.Project+"%")
	}
	//if data.Creator != 0 {
	//	db = db.Where("creator = ?", data.Creator)
	//}
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询项目列表失败")
	return &list
}

func (m *projectModelImpl) Update(data entity.AiderProject) *entity.AiderProject {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改项目失败")
	return &data
}

func (m *projectModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.AiderProject{}, "id in (?)", ids).Error, "删除项目失败")
}
