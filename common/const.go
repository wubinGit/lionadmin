package common

const (
	BASE_URL="api/v1/admin"


	ADMIN_BASE_URL=BASE_URL

	PERMISSION_BASE_URL=BASE_URL+"/permission"

	ROLE_BASE_URL=BASE_URL+"/role"

	MACHINE_BASE_URL=BASE_URL+"/machine"

	APP_BASE_URL=BASE_URL+"/app"

	SSH_BASE_URL=BASE_URL+"/ssh"


	SFTP_BASE_URL=BASE_URL+"/sftp"

	SYNCD_BASE_URL=BASE_URL+"/syncd"


	LOGIN     = "/login"
	REGITSTER = "/register"

	ADMIN_CREATE = "/add"
	ADMIN_UPDATE = "/update"
	ADMIN_GETLIST = "/list"
	ADMIN_GETLIST_EXCEPT = "/list/machine"
	ADMIN_DELETE = "/delete/:id"
	ADMIN_BYPAGE = "/getbypage"

	/**
	adminpermission url
	 */
	PERMISSION_GETALL="/list"
	PERMISSION_ROLE_ID="/getbyid/:id"
	PERMISSION_UPDATE="/update/:id"
	PERMISSION_BYPAGE="/getpage"

	/**
	role
	 */
	PERMISSION_ROLE_ADD="/add"
	PERMISSION_ROLE_UPDATE="/update"
	PERMISSION_ROLE_DELETE="/delete/:id"
	PERMISSION_ROLE_SELECT="/getall"
	PERMISSION_ROLE_PAGE="/getbypage"

	/**
	machine
	 */
	MACHINE_ADD="/add"
	MACHINE_UPDATE="/update"
	MACHINE_DELETE="/delete/:id"
	MACHINE_SELECT="/getbyid/:id"
	MACHINE_SELECT_PAGE="/getbypage"
	MACHINE_SELECT_AUTH="/auth"
	MACHINE_SELECT_ALL="/all"

	/**
	app
	 */
	APP_ADD="/add"
	APP_UPDATE="/update"
	APP_DELETE="/delete/:id"
	APP_SELECT="/getbyid/:id"
	APP_SELECT_PAGE="/getbypage"

	/**
	ssh
	 */
	SSH_CONNECT="/connect/:id"
	/**
	test
	 */
	SSH_TEST_CMD="/cmd"

	/**
	sftp
	 */
	SFTP_FILE="/sftp/:id"
	SFTP_CAT_FILE="/sftp/:id/cat"
	SFTP_RENAME_FILE="/sftp/:id/rename"
	SFTP_MKDIR_FILE="/sftp/:id/mkdir"
	SFTP_RM_FILE="/sftp/:id/rm"
	SFTP_DOWNLOAD_FILE="/sftp/:id/dl"
	SFTP_UPLOAD_FILE="/sftp/:id/up"

	/**
	syncd
	 */
	SYNCD_PAGE_LIST="/syncd/page"
	SYNCD_BUILD="/syncd/build"
	SYNCD_BUILD_STATUS_LOG="/syncd/buildLog/:id"
	SYNCD_FILE_TAR="/syncd/tar"
	SYNCD_CREATE_OR_UPATAE="/syncd/createOrUpdate"
	SYNCD_PROJECT_GETBYID="/syncd/project/:id"
	SYNCD_UPDATE_PROJECT="/syncd/updataProject"
	SYNCD_DEPLOY_PROJECT="/syncd/deploy"
	SYNCD_DEPLOY_PROJECT_CMD="/syncd/deploy/cmd/:id"
	SYNCD_DEPLOY_STATUS_LOG="/syncd/deployLog/:id"


	/**
	win
	 */
	SYNCD_BUILD_WIN="/build"




)
