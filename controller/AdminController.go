package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/model"
	"lionadmin.org/lion/request"
	"lionadmin.org/lion/service"
	_ "lionadmin.org/lion/util"
	"lionadmin.org/lion/util/baseResult"
	"lionadmin.org/lion/util/gostring"
	"strings"
)



/**
创建管理员
*/
func CreateAdmin(c *gin.Context) {
	var admin model.TAdmin
	var requestAdmin request.RequestAdmin

	if err := c.ShouldBind(&requestAdmin); err != nil {
		common.ParamError(c, "参数不完整")
		return
	}

	if len(requestAdmin.Username) == 0 || len(requestAdmin.PasswordHash) == 0 {
		common.ParamError(c, "用户名密码不能为空")
		return
	}
	if requestAdmin.RoleId == 0 {
		common.ParamError(c, "角色不能为空")
		return
	}
	if strings.Compare(requestAdmin.PasswordHash,requestAdmin.ResetPassword)!=0 {
		common.ParamError(c, "密码不一致")
		return
	}

	admin.Username=requestAdmin.Username
	admin.PasswordHash=requestAdmin.PasswordHash
	admin.RoleId=requestAdmin.RoleId
	admin.Status=requestAdmin.Status

	hashPassword, error := bcrypt.GenerateFromPassword([]byte(admin.PasswordHash), bcrypt.DefaultCost)
	if error != nil {
		common.AppError(c, "加密错误")
		return
	}
	newAdmin := model.TAdmin{
		Username:     admin.Username,
		PasswordHash: string(hashPassword),
		RoleId:       admin.RoleId,
	}

	err := service.CreateAdmin(newAdmin)

	baseResult.Create(c, err)
}

/**
删除用户
*/
func DeleteAdmin(c *gin.Context) {
	param := c.Param("id")
	id := gostring.Str2Int(param)
	if id == 0 {
		common.AppError(c, "ID不能为空")
		return
	}
	err := service.DeleteAdmin(id)
	baseResult.Delete(c, err)
}

/**
更新信息
*/
func UpdateAdmin(c *gin.Context) {
	var admin model.TAdmin
	var requestAdmin request.RequestAdmin
	if err := c.ShouldBind(&requestAdmin); err != nil {
		common.ParamError(c, "参数不完整")
		return
	}
	if len(requestAdmin.Username) == 0 || len(requestAdmin.PasswordHash) == 0 {
		common.ParamError(c, "用户名密码不能为空")
		return
	}
	if requestAdmin.RoleId == 0 {
		common.ParamError(c, "角色不能为空")
		return
	}
	if len(string(requestAdmin.ID)) == 0 {
		common.ParamError(c, "ID不能为空")
		return
	}
	if strings.Compare(requestAdmin.PasswordHash,requestAdmin.ResetPassword)!=0 {
		common.ParamError(c, "密码不一致")
		return
	}
	admin.ID=requestAdmin.ID
	admin.Username=requestAdmin.Username

	admin.RoleId=requestAdmin.RoleId
	admin.Status=requestAdmin.Status
	hashPassword, error := bcrypt.GenerateFromPassword([]byte(requestAdmin.PasswordHash), bcrypt.DefaultCost)
	if error != nil {
		common.AppError(c, "加密错误")
		return
	}
	admin.PasswordHash=string(hashPassword)
	err := service.UpdateAdmin(&admin)
	baseResult.Update(c, err)
}

//分页查询
func GetAdminList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	info := request.PageInfoData(pageInfo)
	err, list, total := service.GetAdminInfoList(info)
	baseResult.PageData(c, err, list, total)

}

func GetAdminAll(c *gin.Context)  {
	err, list := service.GetAdminAll()
	baseResult.Select(c,err,list)
}

func GetAdminExceptCurrentUser(c *gin.Context)  {
	machineId := c.Query("id")
	if len(machineId)==0 {
		common.ParamError(c, "ID不能为空")
		return
	}
		err, list := service.GetAdminExceptCurrentUser(gostring.Str2Int(machineId))
		baseResult.Select(c,err,list)


}
