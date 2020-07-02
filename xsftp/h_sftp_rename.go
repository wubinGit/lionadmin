package xsftp

import (
	"github.com/gin-gonic/gin"
	"lionadmin.org/lion/common"
)

func SftpRename(c *gin.Context) {
	sftpClient, error := getSftpClient(c)
	if error!=nil {
		common.AppError(c,"连接错误")
		return
	}
	oPath := c.Query("opath")
	nPath := c.Query("npath")
	err := sftpClient.Rename(oPath, nPath)
	if err!=nil {
		common.AppError(c,"修改错误")
		return
	}
	common.JsonSuccess(c)
}
