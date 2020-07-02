package route

import (
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/controller"
)

func SshRouter(baseUrl string)  {
	r := Router.Group("/" + baseUrl)
	r.GET(common.SSH_CONNECT,controller.SshConnect)
	r.GET(common.SSH_TEST_CMD,controller.SshRuncommend)
}
