package model


type TRemoteBuildLog struct {
	ID                  int     `gorm:"column:id"`
	ApplyId             int     `gorm:"column:apply_id"`
	StartTime           int     `gorm:"column:start_time"`
	FinishTime          int     `gorm:"column:finish_time"`
	Status              int     `gorm:"status"`
	Tar                 string  `gorm:"tar"`
	Output              string  `gorm:"output"`
	Errmsg              string  `gorm:"errmsg"`
	Ctime               int     `gorm:"ctime"`
}

const (
	BUILD_SUCCESS=0
	BUILD_FILE=1
	//正在构建
	BUILD_ON=2

)
