package model

import (
      "lionadmin.org/lion/util/utils"
)

type TMachine struct {
      //gorm.Model
      ID  int `json:"id" gorm:"column:id" form:"id"`
      Name  string `json:"name" gorm:"column:name" form:"name"`
      SshPassword  string `json:"password" gorm:"column:ssh_password" form:"password"`
      Status  int `json:"status" gorm:"column:status" form:"status"`
      UpdatedAt  utils.JSONTime `json:"UpdatedAt" gorm:"column:updated_at"`
      CreatedAt   utils.JSONTime     `json:"createdAt" gorm:"column:created_at"`
      SshIp  string `json:"ip" gorm:"column:ssh_ip" form:"ip"`
      SshPort  string `json:"port" gorm:"column:ssh_port" form:"port"`
      Desc  string `json:"desc" gorm:"column:desc" form:"desc"`
      //AdminId  int `json:"adminId" gorm:"column:admin_id" form:"adminId"`
}