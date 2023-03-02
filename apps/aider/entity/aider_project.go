package entity

import (
	"github.com/XM-GO/PandaKit/model"
	"time"
)

// AiderEvent 事件控制表
type AiderProject struct {
	Project    string    `json:"project" gorm:"type: varchar(255);column:项目"` // 项目名称  <--> 文件名
	Progress   int       `json:"progress" gorm:"type: int;comment:完成进度"`
	LastTask   string    `json:"lastTask" gorm:"type: varchar(255);column:上次完成"`
	ThisTask   string    `json:"thisTask" gorm:"type: varchar(255);column:本次待做"`
	StartTime  time.Time `json:"startTime" gorm:"column:开始时间"`
	FinishTime time.Time `json:"finishTime" gorm:"column:结束时间"`
	Status     string    `json:"status" gorm:"type: varchar(1);column:完成状态"`
	model.BaseAutoModel
}
