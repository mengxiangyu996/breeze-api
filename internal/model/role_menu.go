package model

import (
	"breeze-api/pkg/datetime"

	"gorm.io/gorm"
)

// 角色菜单关系模型
type RoleMenu struct {
	Id         int           `gorm:"autoIncrement"`
	CreateTime datetime.Time `gorm:"autoCreateTime"`
	UpdateTime datetime.Time `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
	RoleId     int
	MenuId     int
}
