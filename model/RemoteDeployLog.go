package model

type TRemoteDeployLog struct {
	ID                  int     `gorm:"column:id"`
	ApplyId             int     `gorm:"column:apply_id"`
	Status              int     `gorm:"status"`
	Content             string  `gorm:"output"`
	Ctime               int     `gorm:"ctime"`
	
}
