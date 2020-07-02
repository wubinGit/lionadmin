package route

import (
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/controller"
)

func AdminRouter(base string)  {
	r := Router.Group("/" + base)
	//r.POST(common.REGITSTER, controller.Register)
	r.POST(common.LOGIN, controller.Login)
	r.POST(common.ADMIN_CREATE, controller.CreateAdmin)
	r.DELETE(common.ADMIN_DELETE, controller.DeleteAdmin)
	r.GET(common.ADMIN_BYPAGE, controller.GetAdminList)
	r.PUT(common.ADMIN_UPDATE, controller.UpdateAdmin)
	r.GET(common.ADMIN_GETLIST, controller.GetAdminAll)
	r.GET(common.ADMIN_GETLIST_EXCEPT, controller.GetAdminExceptCurrentUser)
}
