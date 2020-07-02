package route

import (
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/controller"
)

func SyncdRouter(baseUrl string)  {
	group := Router.Group("/"+baseUrl)
	/**
	获取项目列表
	 */
	group.GET(common.SYNCD_PAGE_LIST,controller.GetProjiectList)
	/**
	构建项目
	 */
	group.GET(common.SYNCD_BUILD,controller.BuildProject)
	group.GET(common.SYNCD_BUILD_STATUS_LOG,controller.BuildStatus)
	//group.GET(common.SYNCD_FILE_TAR,controller.GetFileTar)
	/**
	创建项目
	 */
	group.POST(common.SYNCD_CREATE_OR_UPATAE,controller.ProjectCreateOrUpdate)
	/**

	 */
	group.GET(common.SYNCD_PROJECT_GETBYID,controller.GetProjectById)
	/**
	设置项目打包执行的脚本
	 */
	group.PUT(common.SYNCD_UPDATE_PROJECT,controller.UpdateProject)


	group.GET(common.SYNCD_DEPLOY_PROJECT,controller.DeployStart)

	group.GET(common.SYNCD_BUILD_WIN,controller.WinBuild)

	group.GET(common.SYNCD_DEPLOY_PROJECT_CMD,controller.DeployCmdExe)

	group.GET(common.SYNCD_DEPLOY_STATUS_LOG,controller.DepleyStatus)







}
