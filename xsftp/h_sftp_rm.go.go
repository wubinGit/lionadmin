package xsftp

import (
	"github.com/gin-gonic/gin"
	"lionadmin.org/lion/common"
)

func SftpRm(c *gin.Context) {
	sftpClient, _ := getSftpClient(c)
	//if handleError(c, err) {
	//	return
	//}
	fullPath := c.Query("path")
	if fullPath == "/" || fullPath == "$HOME" {
		common.AppError(c,"can't delete / or $HOME dir")
		return
	}

	err := sftpClient.Remove(fullPath)
	if err!=nil {
		common.AppError(c," delete file error")
		return
	}
	common.JsonSuccess(c)
}
