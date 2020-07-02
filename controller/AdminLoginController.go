package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/model"
)

//绑定参数
type loginAdmin struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

/*
jwt
*/
func Login(c *gin.Context) {
	db := common.GetDB()
	//获取参数
	var loginParam = &loginAdmin{}
	//绑定前端参数
	if err := c.ShouldBind(&loginParam); err != nil {
		common.ParamError(c, err.Error())
		return
	}

	//数据验证
	if len(loginParam.Username) == 0 || len(loginParam.Password) == 0 {
		common.ParamError(c, "用户名密码不能为空")
		return
	}
	/**
	数据库实体
	*/
	var admin model.TAdmin
	db.Where("username=?", loginParam.Username).Find(&admin)
	//用户不存在
	if admin.ID == 0 {
		common.NoDataError(c, "用户不存在")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(loginParam.Password));
		err != nil {
		common.ParamError(c, "密码错误")
		return
	}

	//发送token
	token, err := common.ReleaseToken(admin)
	if err != nil {
		common.AppError(c, "系统异常")
		return
	}
	//返回结果
	common.JSON(c, token)
}
