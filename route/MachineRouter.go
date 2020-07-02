package route

import (
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/controller"
)

func MachineRouter(baseUrl string)  {
	r := Router.Group("/" + baseUrl)
	r.POST(common.MACHINE_ADD,controller.CreateMachine)
	r.PUT(common.MACHINE_UPDATE,controller.UpdateMachine)
	r.DELETE(common.MACHINE_DELETE,controller.DeleteMachine)
	r.GET(common.MACHINE_SELECT,controller.FindMachine)
	r.GET(common.MACHINE_SELECT_PAGE,controller.GetMachineList)
	r.GET(common.MACHINE_SELECT_ALL,controller.GetManchines)
	r.POST(common.MACHINE_SELECT_AUTH,controller.AuthorizationUser)

}
