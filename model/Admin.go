package model

import (
	"lionadmin.org/lion/util/utils"
)

type TAdmin struct {
	//gorm.Model
	Username     string `json:"username" gorm:"column:username" form:"username" `
	PasswordHash string `json:"passwordHash" gorm:"column:password_hash" form:"password" `
	RoleId       int    `json:"roleId" gorm:"column:role_id" form:"roleId"`
	CreateAt     utils.JSONTime     `json:"createAt" gorm:"column:created_at"`
	UpdateAt     utils.JSONTime      `json:"updateAt" gorm:"column:updated_at"`
	Status       int    `json:"status" gorm:"column:status" form:"status"`
	ID           int    `json:"id" gorm:"column:id" form:"id"`
}
