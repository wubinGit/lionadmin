package model

type TApp struct {
     // gorm.Model
      ID  int `json:"id" gorm:"column:id" form:"id"`
      Name  string `json:"name" gorm:"column:name" form:"name"`
      Status  int `json:"status" gorm:"column:status" form:"status"`
      Desc  string `json:"desc" gorm:"column:desc" form:"desc"`
      Type  int `json:"type" gorm:"column:type" form:"type"`
}