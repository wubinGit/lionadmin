package xsftp

import (
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/model"
	"lionadmin.org/lion/service"
	"lionadmin.org/lion/ssh"
	"log"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/pkg/sftp"
)

const maxPacket = 1 << 15

func SftpMkdir(c *gin.Context) {
	sftpClient, _ := getSftpClient(c)
	//if handleError(c, err) {
	//	return
	//}
	fullPath := c.Query("path")
	if strings.HasPrefix(fullPath, "$HOME") {
		wd, _ := sftpClient.Getwd()
		//if handleError(c, err) {
		//	return
		//}
		fullPath = strings.Replace(fullPath, "$HOME", wd, 1)
	}
	log.Println(fullPath)

	_ = sftpClient.Mkdir(fullPath)
	//if handleError(c, err) {
	//	returnc
	//}
	common.JsonSuccess(c)
}

func getSftpClient(c *gin.Context) (*sftp.Client, error) {
	idx, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, err
	}
	err, machine := service.GetMachine(idx)
	if err != nil {
		return nil, err
	}
	client, err := NewSftpClient(machine)
	return client, nil

}




func NewSftpClient(h  model.TMachine) (*sftp.Client, error) {
	conn, err := ssh.NewSshClient(h)
	if err != nil {
		return nil, err
	}
	return sftp.NewClient(conn, sftp.MaxPacket(maxPacket))
}
