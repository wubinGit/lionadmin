package controller

import (
	"github.com/gin-gonic/gin"
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/log"
	"lionadmin.org/lion/model"
	"lionadmin.org/lion/request"
	"lionadmin.org/lion/service"
	"lionadmin.org/lion/util/baseResult"
	"lionadmin.org/lion/util/gostring"
	"lionadmin.org/lion/util/utils"
	"net"
)

func CreateMachine(c *gin.Context) {
	var machine model.TMachine
	_ = c.ShouldBind(&machine)

	encrypt, done := PasswordEnconde(c, machine)
	if !done {
		return
	}
	machine.SshPassword = encrypt
	if !checkIpRule(c, machine.SshIp) {
		return
	}
	err := service.CreateMachine(machine)
	baseResult.Create(c, err)
}

func DeleteMachine(c *gin.Context) {
	id := gostring.Str2Int(c.Param("id"))
	err := service.DeleteMachine(id)
	baseResult.Delete(c, err)
}

func UpdateMachine(c *gin.Context) {
	var machine model.TMachine
	_ = c.ShouldBind(&machine)
	if machine.ID == 0 {
		common.ParamError(c, "ID不能为空")
		return
	}
	encrypt, done := PasswordEnconde(c, machine)
	if !done {
		return
	}
	machine.SshPassword = encrypt
	if !checkIpRule(c, machine.SshIp) {
		return
	}
	err := service.UpdateMachine(&machine)
	baseResult.Update(c, err)

}

/**
给用户授权登录ssh
 */


func AuthorizationUser(c *gin.Context)  {
	var machineUser request.TMachineUsersRequest
	err := c.ShouldBind(&machineUser)
	if err!=nil {
		common.ParamError(c, "绑定参数错误")
		return
	}
	if machineUser.MachineId==0 || len(machineUser.AdminId)==0 {
		common.ParamError(c, "参数错误")
		return
	}
	err = service.AuthorizationUser(machineUser)
	baseResult.Auth(c, err)

}




/**
密码加密
*/
func PasswordEnconde(c *gin.Context, machine model.TMachine) (string, bool) {
	if len(machine.Name) == 0 {
		common.ParamError(c, "参数有误")
		return "", false
	}
	if len(machine.SshPassword) == 0 {
		common.ParamError(c, "密码不能为空")
		return "", false
	}
	//TODO 第二个参数为秘钥 用来加密解密过程  这里先固定  一定要8位长度
	encrypt, encrypterr := utils.Encrypt(machine.SshPassword, []byte("xlf12345"))
	if encrypterr != nil {
		common.AppError(c, "系统错误")
		log.Error("密码加密错误")
		return "", false
	}
	return encrypt, true
}

func FindMachine(c *gin.Context) {
	param := c.Param("id")
	id := gostring.Str2Int(param)
	err, remachine := service.GetMachine(id)
	baseResult.Select(c, err, remachine)
}

func GetMachineList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}
	err, list, total := service.GetMachineInfoList(pageInfo)
	baseResult.PageData(c, err, list, total)
}

func checkIpRule(c *gin.Context, ip string) (result bool) {
	if len(ip) == 0 {
		common.ParamError(c, "IP规则不能为空")
		return false
	}
	ipAddress := net.ParseIP(ip)
	if ipAddress == nil {
		common.ParamError(c, "IP不符合规则")
		return false
	}
	return true
}

func GetManchines(c *gin.Context)  {
	err, list, total := service.GetManchines()
	baseResult.PageData(c, err, list, total)

}



