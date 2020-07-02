package controller

import (
	"github.com/gin-gonic/gin"
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/model"
	"lionadmin.org/lion/request"
	"lionadmin.org/lion/service"
	"lionadmin.org/lion/util/baseResult"
	"lionadmin.org/lion/util/gostring"
)



func FindAdminPermission(c *gin.Context) {
	var permission model.TAdminPermission
	_ = c.ShouldBindQuery(&permission)
	err, repermission := service.GetAdminPermission(permission.ID)
	baseResult.Select(c,err,repermission)
}

/**
查询所有
*/
func GetAdminPermissionAll(c *gin.Context) {
	err, list := service.GetPermissionAll()
	baseResult.Select(c,err,list)

}

/**
查询角色所有的权限
*/
func GetPermissionByRoleId(c *gin.Context) {
	id := gostring.Str2Int(c.Param("id"))
	if id == 0 {
		common.ParamError(c, "角色ID不能为空")
		return
	}
	err, list := service.GetPermissionByRoleId(id)
	baseResult.Select(c,err,list)
}

/**
更新角色的权限
*/
func UpdatePermissionRole(c *gin.Context) {
	var permmission request.Permissions
	param := c.Param("id")
	if len(param)==0 {
		common.ParamError(c,"参数错误")
		return
	}
	if  err := c.ShouldBind(&permmission);err!=nil{
		common.ParamError(c,"参数错误")
		return
	}
	if permmission.Permissions==nil {
		common.ParamError(c,"参数错误")
		return
	}
	err:= service.UpdatePermissionRole(param, permmission.Permissions)
	baseResult.Update(c,err)

}

/**
page
*/
func GetAdminPermissionList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if pageInfo.Page==0 {
		pageInfo.Page=1
	}
	if pageInfo.PageSize==0 {
		pageInfo.PageSize=10
	}
	err, list, total := service.GetAdminPermissionInfoList(pageInfo)
	baseResult.PageData(c,err,list,total)
}
