package initialize

import (
	devEntity "GinMasterReWi/apps/develop/entity"
	flowEntity "GinMasterReWi/apps/flow/entity"
	jobEntity "GinMasterReWi/apps/job/entity"
	logEntity "GinMasterReWi/apps/log/entity"
	resSourceEntity "GinMasterReWi/apps/resource/entity"
	"GinMasterReWi/apps/system/entity"
	"GinMasterReWi/pkg/global"
	"github.com/XM-GO/PandaKit/biz"
)

// 初始化时如果没有表创建表
func InitTable() {
	m := global.Conf.Server
	if m.IsInitTable {
		biz.ErrIsNil(
			global.Db.AutoMigrate(
				//casbin.CasbinRule{},
				entity.SysDept{},
				entity.SysApi{},
				entity.SysConfig{},
				entity.SysDictType{},
				entity.SysDictData{},
				logEntity.LogLogin{},
				logEntity.LogOper{},
				logEntity.LogJob{},
				entity.SysUser{},
				entity.SysTenants{},
				entity.SysRole{},
				entity.SysMenu{},
				entity.SysPost{},
				entity.SysRoleMenu{},
				entity.SysRoleDept{},
				entity.SysNotice{},
				jobEntity.SysJob{},
				devEntity.DevGenTable{},
				devEntity.DevGenTableColumn{},
				resSourceEntity.ResOss{},
				resSourceEntity.ResEmail{},

				flowEntity.FlowWorkClassify{},
				flowEntity.FlowWorkInfo{},
				flowEntity.FlowWorkTemplates{},
				flowEntity.FlowWorkOrder{},
				flowEntity.FlowWorkOrderTemplate{},
				flowEntity.FlowWorkStage{},
				flowEntity.FlowWorkTask{},
				flowEntity.FlowWorkTaskHistory{},
			),
			"初始化表失败")
	}
}
