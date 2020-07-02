package route

import (
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/controller"
)

func AppRouter(baseUrl string)  {
	r := Router.Group("/" + baseUrl)
	r.POST(common.APP_ADD,controller.CreateApp)
	r.PUT(common.APP_UPDATE,controller.UpdateApp)
	r.DELETE(common.APP_DELETE,controller.DeleteApp)
	r.GET(common.APP_SELECT,controller.FindApp)
	r.GET(common.APP_SELECT_PAGE,controller.GetAppList)
}
