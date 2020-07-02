package service

import (
	"fmt"
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/model"
	"lionadmin.org/lion/request"
)

func CreateMachine(machine model.TMachine) (err error) {
	db := common.GetDB()
	err = db.Create(&machine).Error
	return err
}

func DeleteMachine(id int) (err error) {
	db := common.GetDB()
	err = db.Where("id = ?", id).Delete(&model.TMachine{}).Error
	return err
}

func UpdateMachine(machine *model.TMachine) (err error) {
	db := common.GetDB()
	err = db.Model(&machine).Updates(model.TMachine{Name: machine.Name, Status: machine.Status, SshIp: machine.SshIp,
		SshPassword: machine.SshPassword,SshPort:machine.SshPort, Desc: machine.Desc}).Error
	return err
}

func GetMachine(id int) (err error, machine model.TMachine) {
	db := common.GetDB()
	err = db.Where("id = ?", id).First(&machine).Error
	return err, machine
}

func GetMachineInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := common.GetDB()
	var machines []model.TMachine
	err = db.Find(&machines).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&machines).Error
	return err, machines, total
}

/**
判断当前用户是否有权限登录该机器
*/
func JudgeUserIsPermission(machineId int, admin interface{}) (result bool,machineResult model.TMachine,err error) {
	db := common.GetDB()
	var machine model.TMachine
	//接口类型转换
	adminUser:=admin.(model.TAdmin)
	var machineUsers []model.TMachineUsers
	found := db.Where("id = ?", machineId).First(&machine).RecordNotFound()
	if !found {
		err := db.Where("machine_id = ? ", machineId).Find(&machineUsers).Error
		if err == nil {
			for _, machineUser := range machineUsers {
				if machineUser.UserId == adminUser.ID {
					return true,machine, err
				}
			}
		}
	}
	return false, machine,err
}

/**
授权用户ssh登录
 */
func AuthorizationUser(machine request.TMachineUsersRequest) (err error) {
	sql := "INSERT INTO `t_machine_users` (`user_id`,`machine_id`) VALUES "
	db := common.GetDB()
	// 循环data数组,组合sql语句
	for key, value := range machine.AdminId {
		if len(machine.AdminId)-1 == key {
			//最后一条数据 以分号结尾
			sql += fmt.Sprintf("(%d,'%d');", value, machine.MachineId)
		} else {
			sql += fmt.Sprintf("(%d,'%d'),", value, machine.MachineId)
		}
	}
	db.Exec(sql)
	return nil

}

func GetManchines()(err error,list interface{},total int) {
	db := common.GetDB()
	var manchines []model.TMachine
	err = db.Find(&manchines).Error
	err = db.Find(&manchines).Count(&total).Error
	return err,manchines,total
}
