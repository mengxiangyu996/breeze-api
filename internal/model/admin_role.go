package model

import (
	"breeze-api/pkg/datetime"

	"gorm.io/gorm"
)

// 管理员角色关系模型
type AdminRole struct {
	Id         int           `gorm:"autoIncrement"`
	CreateTime datetime.Time `gorm:"autoCreateTime"`
	UpdateTime datetime.Time `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
	AdminId    int
	RoleId     int
}
