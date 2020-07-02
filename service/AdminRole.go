package service

import (
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/model"
	"lionadmin.org/lion/request"
)

func CreateAdminRole(role model.TAdminRole) (err error) {
	db := common.GetDB()
	err = db.Create(&role).Error
	return err
}

func DeleteAdminRole(id int) (err error) {
	db := common.GetDB()
	err = db.Where("id = ?", id).Delete(model.TAdminRole{}).Error
	return err
}

/**
更新角色
 */
func UpdateAdminRole(role *model.TAdminRole) (err error) {
	db := common.GetDB()
	err = db.Model(&role).Updates(model.TAdminRole{ RoleDesc: role.RoleDesc}).Error
	return err
}

/**
查询一条记录
*/
func GetAdminRole(id uint) (err error, role model.TAdminRole) {
	db := common.GetDB()
	err = db.Where("id = ?", id).First(&role).Error
	return
}

func GetAdminRoleInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := common.GetDB()
	var roles []model.TAdminRole
	err = db.Find(&roles).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&roles).Error
	return err, roles, total
}

/**
获取所有的权限
 */
func GetAdminRoles()(err error,list interface{},total int) {
	db := common.GetDB()
	var roles []model.TAdminRole
	err = db.Find(&roles).Error
	err = db.Find(&roles).Count(&total).Error
	return err,roles,total
}



