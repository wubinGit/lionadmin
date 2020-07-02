package service

import (
	"fmt"
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/model"
	"lionadmin.org/lion/request"
)

func GetAdminPermission(id int) (err error, permission model.TAdminPermission) {
	db := common.GetDB()
	err = db.Where("id = ?", id).First(&permission).Error
	return
}

/**
查询所有权限
*/
func GetPermissionAll() (err error, list interface{}) {
	db := common.GetDB()
	var permisssions []model.TAdminPermission
	err = db.Find(&permisssions).Error
	return err, permisssions
}

/**
查询角色所有的权限
*/
func GetPermissionByRoleId(roleId int) (err error, permissions interface{}) {
	db := common.GetDB()
	var adminpermission []model.TAdminPermission
	//原生sql 查询
	err = db.Raw("select id, name, description, created_at, updated_at from t_admin_permission where id in (select permission_id from t_admin_role_permission where role_id = ?)", roleId).Scan(&adminpermission).Error
	return err, adminpermission
}

/**
更新用户角色
*/
func UpdatePermissionRole(roleId string, permissions []int) (err error) {
	/**
	删除已有的权限
	*/
	db := common.GetDB()

	err = db.Exec("delete from t_admin_role_permission where role_id = ?", roleId).Error
	/**
	保存当前的权限
	*/
	sql := "INSERT INTO t_admin_role_permission ( role_id, permission_id ) VALUES "
	for key, value := range permissions {
		if len(permissions)-1 == key {
			sql += fmt.Sprintf("(%s,%d);", roleId, value)
		} else {
			sql += fmt.Sprintf("(%s,%d),", roleId, value)
		}
	}
	err = db.Exec(sql).Error
	return err

}

/**
分页查询
*/
func GetAdminPermissionInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := common.GetDB()
	var permissions []model.TAdminPermission
	err = db.Find(&permissions).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&permissions).Error
	return err, permissions, total
}
