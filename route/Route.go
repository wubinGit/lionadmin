package route

import (
	"github.com/gin-gonic/gin"
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/middleware"
)

var Router *gin.Engine

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware()) //配置跨域

	r.Use(middleware.AuthMiddleware()) //登录拦截

	Router = r


	AdminRouter(common.ADMIN_BASE_URL)
	AdminPermissionRouter(common.PERMISSION_BASE_URL)
	AdminRoleRouter(common.ROLE_BASE_URL)
	MachineRouter(common.MACHINE_BASE_URL)
	AppRouter(common.APP_BASE_URL)
	SshRouter(common.SSH_BASE_URL)
	SftpRouter(common.SFTP_BASE_URL)
	SyncdRouter(common.SYNCD_BASE_URL)
	return r
}
