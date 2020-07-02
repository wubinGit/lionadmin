package route

import (
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/controller"
)

func AdminRoleRouter(baseUrl string)  {

	r := Router.Group("/" + baseUrl)

	r.POST(common.PERMISSION_ROLE_ADD,controller.CreateAdminRole)
	r.PUT(common.PERMISSION_ROLE_UPDATE,controller.UpdateAdminRole)
	r.DELETE(common.PERMISSION_ROLE_DELETE,controller.DeleteAdminRole)
	r.GET(common.PERMISSION_ROLE_SELECT,controller.GetAdminRoles)
	r.GET(common.PERMISSION_ROLE_PAGE,controller.GetTAdminRoleList)

}
