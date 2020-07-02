package main

import (
	"github.com/gin-gonic/gin"
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/controller"
	"lionadmin.org/lion/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	//跨域
	r.Use(middleware.CORSMiddleware())
	//token 校验 测试环境不需要
	//r.Use( middleware.AuthMiddleware())
	//请求
	//r.POST(common.REGITSTER, controller.Register)
	r.POST(common.LOGIN, controller.Login)
	r.POST(common.ADMIN_CREATE, controller.CreateAdmin)
	r.DELETE(common.ADMIN_DELETE, controller.DeleteAdmin)
	r.GET(common.ADMIN_GETLIST, controller.GetAdminList)
	r.PUT(common.ADMIN_UPDATE, controller.UpdateAdmin)
	//r.GET(common.USER_INFO, middleware.AuthMiddleware(), controller.Info)
	return r
}
