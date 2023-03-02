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
	AiderDateModel interface {
		Insert(data entity.AiderDate) *entity.AiderDate
		FindOne(id int64) *entity.AiderDate
		FindListPage(page, pageSize int, data entity.AiderDate) (*[]entity.AiderDate, int64)
		FindList(data entity.AiderDate) *[]entity.AiderDate
		Update(data entity.AiderDate) *entity.AiderDate
		Delete(ids []int64)
	}

	dateModelImpl struct {
		table string
	}
)

var AiderDateModelDao AiderDateModel = &dateModelImpl{
	table: `aider_events`,
}

func (m *dateModelImpl) Insert(data entity.AiderDate) *entity.AiderDate {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加日程控制失败")
	return &data
}

func (m *dateModelImpl) FindOne(id int64) *entity.AiderDate {
	resData := new(entity.AiderDate)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询日程控制失败")
	return resData
}

func (m *dateModelImpl) FindListPage(page, pageSize int, data entity.AiderDate) (*[]entity.AiderDate, int64) {
	list := make([]entity.AiderDate, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	//if data.Event01 != "" {
	//	db = db.Where("name like ?", "%"+data.Event01+"%")
	//}
	//if data.Creator != 0 {
	//	db = db.Where("creator = ?", data.Creator)
	//}
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询日程控制分页列表失败")
	return &list, total
}

func (m *dateModelImpl) FindList(data entity.AiderDate) *[]entity.AiderDate {
	list := make([]entity.AiderDate, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	//if data.Event01 != "" {
	//	db = db.Where("name like ?", "%"+data.Event01+"%")
	//}
	//if data.Creator != 0 {
	//	db = db.Where("creator = ?", data.Creator)
	//}
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询日程控制列表失败")
	return &list
}

func (m *dateModelImpl) Update(data entity.AiderDate) *entity.AiderDate {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改日程控制失败")
	return &data
}

func (m *dateModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.AiderDate{}, "id in (?)", ids).Error, "删除日程控制失败")
}
