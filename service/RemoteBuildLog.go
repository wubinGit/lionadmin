package service

import (
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/model"
)

func CreateBuildLog(buildlog *model.TRemoteBuildLog) (err error)  {
	db := common.GetDB()
	err = db.Create(buildlog).Error
	return err
}

func SelectBuildLogStatus(id int)(result model.TRemoteBuildLog,err error)  {
	var BuildLogStatus model.TRemoteBuildLog
	db := common.GetDB()
	err = db.Where("apply_id = ?", id).Last(&BuildLogStatus).Error
	return BuildLogStatus,err

}

func UpdateBuildLogStatus(BuildLogStatus *model.TRemoteBuildLog)(err error)  {
	db := common.GetDB()
	db.Model(&BuildLogStatus).Updates(model.TRemoteBuildLog{Status: BuildLogStatus.Status,Output:BuildLogStatus.Output,FinishTime: BuildLogStatus.FinishTime})
	return err

}


