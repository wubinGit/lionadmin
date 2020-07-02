package model

import "lionadmin.org/lion/util/utils"

type TRemoteProject struct {
	ID                  int     `gorm:"primary_key"`
	Name                string  `gorm:"type:varchar(100);not null;default:''"`
	Description         string  `gorm:"type:varchar(500);not null;default:''"`
	NeedAudit           int     `gorm:"type:int(11);not null;default:0"`
	Status		        int	`gorm:"type:int(11);not null;default:0"`
	RepoUrl             string  `gorm:"type:varchar(500);not null;default:''"`
	DeployMode		    int	`gorm:"type:int(11);not null;default:0"`
	RepoBranch          string  `gorm:"type:varchar(100);not null;default:''"`
	OnlineCluster       string  `gorm:"type:varchar(1000);not null;default:''"`
	DeployUser          string  `gorm:"type:varchar(100);not null;default:''"`
	DeployPath          string  `gorm:"type:varchar(500);not null;default:''"`
	BuildScript         string  `gorm:"type:text;not null"`
	BuildHookScript     string  `gorm:"type:text;not null"`
	DeployHookScript    string  `gorm:"type:text;not null"`
	PreDeployCmd        string  `gorm:"type:text;not null"`
	AfterDeployCmd      string  `gorm:"type:text;not null"`
	AuditNotice         string  `gorm:"type:varchar(2000);not null;default:''"`
	DeployNotice        string  `gorm:"type:varchar(2000);not null;default:''"`
	CreatedAt  utils.JSONTime  `json:"CreatedAt" gorm:"column:created_at"`
	ManchineId              int   `gorm:"type:int(11);not null;default:1"`
}







