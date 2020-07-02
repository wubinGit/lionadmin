package route

import (
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/controller"
)

func AdminPermissionRouter(baseUrl string)  {
	r := Router.Group("/" + baseUrl)
	r.GET(common.PERMISSION_GETALL,controller.GetAdminPermissionAll)
	r.GET(common.PERMISSION_ROLE_ID,controller.GetPermissionByRoleId)
	r.PUT(common.PERMISSION_UPDATE,controller.UpdatePermissionRole)
	r.GET(common.PERMISSION_BYPAGE,controller.GetAdminPermissionList)

}
