package xsftp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/sftp"
	"io/ioutil"
	"lionadmin.org/lion/common"
	"net/http"
	"os"
)

func ScpFetchFile(c *gin.Context) (*sftp.File, os.FileInfo, error) {
	sftpClient, err := getSftpClient(c)
	if err != nil {
		return nil, nil, err
	}
	fullPath := c.Query("path")
	fileInfo, err := sftpClient.Stat(fullPath)
	if err != nil {
		return nil, nil, err
	}
	if fileInfo.IsDir() {
		return nil, nil, fmt.Errorf("%s is not a file", fullPath)
	}
	f, err := sftpClient.Open(fullPath)
	return f, fileInfo, err
}
func SftpDl(c *gin.Context) {
	file, fileInfo, _ := ScpFetchFile(c)
	defer file.Close()
	//if handleError(c, err) {
	//	return
	//}
	extraHeaders := map[string]string{
		"Content-Disposition": fmt.Sprintf(`attachment; filename="%s"`, fileInfo.Name()),
	}
	c.DataFromReader(http.StatusOK, fileInfo.Size(), "application/octet-stream", file, extraHeaders)
}
func SftpCat(c *gin.Context) {
	file, fileInfo, _ := ScpFetchFile(c)
	if file==nil {
		common.AppError(c,"文件不可读")
		return
	}
	defer file.Close()
	//if handleError(c, err) {
	//	return
	//}
	b, err := ioutil.ReadAll(file)
	if err!=nil {
		common.AppError(c,"系统错误")
		return
	}
	//if handleError(c, err) {
	//	return
	//}
	//c.String(200,"utf-8",file)
	c.AbortWithStatusJSON(200, gin.H{"code": 200, "data": string(b), "msg": fileInfo.Name()})
}
