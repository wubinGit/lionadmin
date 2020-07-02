package model

import (
      "lionadmin.org/lion/util/utils"
)

type TAdminRolePermission struct {
     // gorm.Model
      RoleId  int `json:"roleId" gorm:"column:role_id" form:"roleId"`
      PermissionId  int `json:"permissionId" gorm:"column:permission_id"`
      CreatedAt  utils.JSONTime  `json:"CreatedAt" gorm:"column:created_at"`
      UpdateAt  utils.JSONTime  `json:"UpdateAt" gorm:"column:update_at"`
}