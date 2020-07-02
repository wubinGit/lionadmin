package service

import (
	"fmt"
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/model"
	"lionadmin.org/lion/request"
)

func CreateAdmin(admin model.TAdmin) (err error) {
	db := common.GetDB()
	err = db.Create(&admin).Error
	return err
}

func DeleteAdmin(id int) (err error) {
	db := common.GetDB()
	err = db.Where("id = ?", id).Delete(model.TAdmin{}).Error
	return err
}

func UpdateAdmin(admin *model.TAdmin) (err error) {
	db := common.GetDB()
	err = db.Model(&admin).Updates(model.TAdmin{RoleId: admin.RoleId,PasswordHash: admin.PasswordHash,Username: admin.Username, Status: admin.Status}).Error
	return err
}

func GetAdmin(id uint) (err error, admin model.TAdmin) {
	db := common.GetDB()
	err = db.Where("id = ?", id).First(&admin).Error
	return
}

//分页查询
func GetAdminInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := common.GetDB()
	var admins []model.TAdmin
	if len(info.Search)!=0 {
		db=db.Where("username = ? ",info.Search)
	}
	err = db.Find(&admins).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&admins).Error
	return err, admins, total
}

/**
获取所有的管理员
*/
func GetAdminAll() (err error, list interface{}) {
	db := common.GetDB()
	var adminList []model.TAdmin
	err = db.Find(&adminList).Error
	return err, adminList
}

/**
获取所有用户
*/
func GetAdminExceptCurrentUser(machineId int) (err error, list interface{}) {
	/**
	 */
	var machineList []model.TMachineUsers
	db := common.GetDB()
	err = db.Where("machine_id = ? ", machineId).Find(&machineList).Error
	var adminList []model.TAdmin
	if len(machineList) == 0 {
		err = db.Find(&adminList).Error
		return err, adminList
	}
	sql := " SELECT * FROM `t_admin` WHERE `id` NOT IN ( "
	for key, machine := range machineList {
		if len(machineList)-1 == key {
			sql += fmt.Sprintf("%d);", machine.UserId)
		} else {
			sql += fmt.Sprintf("%d,", machine.UserId)
		}
	}
	db.Raw(sql).Scan(&adminList)
	return err, adminList

}
