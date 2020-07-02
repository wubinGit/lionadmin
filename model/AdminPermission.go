package model

import (
      "lionadmin.org/lion/util/utils"
)

type TAdminPermission struct {
     // gorm.Model
      ID  int `json:"id" gorm:"column:id"`
      Name  string `json:"name" gorm:"column:name"`
      Descriptioin  string `json:"description" gorm:"column:description"`
      CreatedAt  utils.JSONTime  `json:"CreatedAt" gorm:"column:created_at"`
      UpdateAt  utils.JSONTime  `json:"UpdatedAt" gorm:"column:updated_at"`

}