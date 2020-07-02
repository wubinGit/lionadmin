package model

type TMachineUsers struct {
	ID  int `json:"id" gorm:"column:id" form:"id"`
	UserId  int `json:"userId" gorm:"column:user_id" form:"userId"`
	MachineId  int `json:"machineId" gorm:"column:machine_id" form:"machineId"`
	Desc  string `json:"desc" gorm:"column:desc" form:"desc"`
}
