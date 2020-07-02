package controller

import (
	"github.com/gin-gonic/gin"
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/service"
	"lionadmin.org/lion/ssh"
	"lionadmin.org/lion/util/gostring"
)

/**
创建SSH连接
*/
func SshConnect(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		common.AppError(c, "错误参数")
		return
	}
	//鉴权
	//查询数据库 获取ssh机器登录的用户名和密码
	//var admin model.User
	//获取当前用户信息
	user, exists := c.Get("user")
	if exists {
		permission, manchine, err := service.JudgeUserIsPermission(gostring.Str2Int(id), user)
		if err != nil {
			common.AppError(c, "系统错误")
			return
		}
		if !permission {
			common.AppError(c, "权限不足")
			return
		}

		ssh.WsSsh(c, manchine)

	}
	//_, machine := service.GetMachine(gostring.Str2Int(id))
	//
	//ssh.WsSsh(c, machine)

}


