package entity

import (
	"github.com/XM-GO/PandaKit/model"
)

// AiderDate 日程控制表
type AiderDate struct {
	model.BaseAutoModel
	Event01 int `json:"event01" gorm:"type:int;comment:事件一"`
	Event02 int `json:"event02" gorm:"type:int;comment:事件二"`
	Event03 int `json:"event03" gorm:"type:int;comment:事件三"`
	Event04 int `json:"event04" gorm:"type:int;comment:事件四"`
	Event05 int `json:"event05" gorm:"type:int;comment:事件五"`
}
