package service

import (
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/model"
)

func CreateDeployLog(Deploylog *model.TRemoteDeployLog) (err error)  {
	db := common.GetDB()
	err = db.Create(Deploylog).Error
	return err
}

func SelectDeployLogStatus(id int)(result model.TRemoteDeployLog,err error)  {
	var DeployLogStatus model.TRemoteDeployLog
	db := common.GetDB()
	err = db.Where("apply_id = ?", id).Last(&DeployLogStatus).Error
	return DeployLogStatus,err

}

func UpdateDeployLogStatus(DeployLogStatus *model.TRemoteDeployLog)(err error)  {
	db := common.GetDB()
	db.Model(&DeployLogStatus).Updates(model.TRemoteDeployLog{Status: DeployLogStatus.Status,Content:DeployLogStatus.Content})
	return err

}