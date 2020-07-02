package model

import (
      "lionadmin.org/lion/util/utils"
)

type TAdminRole struct {
     // gorm.Model
      ID  int `json:"id" gorm:"column:id" form:"id"`
      RoleName  string `json:"RoleName" gorm:"column:role_name" form:"roleName"`
      RoleDesc  string `json:"roledesc" gorm:"column:role_desc" form:"roleDesc"`
      CreatedAt  utils.JSONTime  `json:"CreatedAt" gorm:"column:created_at"`
      UpdateAt  utils.JSONTime  `json:"UpdateAt" gorm:"column:updated_at"`
}