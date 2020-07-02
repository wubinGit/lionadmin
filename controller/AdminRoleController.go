package controller

import (
	"github.com/gin-gonic/gin"
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/model"
	"lionadmin.org/lion/request"
	"lionadmin.org/lion/service"
	"lionadmin.org/lion/util/baseResult"
	"strconv"
)

func CreateAdminRole(c *gin.Context) {
	var role model.TAdminRole
	_ = c.ShouldBind(&role)
	if len(role.RoleName) == 0 || len(role.RoleDesc) == 0 {
		common.ParamError(c, "参数不能为空")
		return
	}
	err := service.CreateAdminRole(role)
	baseResult.Create(c, err)
}

func DeleteAdminRole(c *gin.Context) {
	param := c.Param("id")
	if len(param) == 0 {
		common.AppError(c, "ID不能为空")
		return
	}
	id, err2 := strconv.Atoi(param)
	if err2 != nil {
		common.AppError(c, "参数错误")
		return
	}
	err := service.DeleteAdminRole(id)
	baseResult.Delete(c, err)
}

func UpdateAdminRole(c *gin.Context) {
	var role model.TAdminRole
	_ = c.ShouldBind(&role)
	err := service.UpdateAdminRole(&role)
	baseResult.Update(c, err)
}

func GetTAdminRoleList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	info := request.PageInfoData(pageInfo)
	err, list, total := service.GetAdminRoleInfoList(info)
	baseResult.PageData(c, err, list, total)


}

func GetAdminRoles(c *gin.Context)  {
	err, list,total := service.GetAdminRoles()
	baseResult.PageData(c, err, list, total)
}
