package xsftp

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/sftp"
	"lionadmin.org/lion/common"
	"mime/multipart"
	"path"
)

func SftpUp(c *gin.Context) {
	sftpClient, _ := getSftpClient(c)
	//if handleError(c, err) {
	//	return
	//}
	file, _ := c.FormFile("file")
	//if handleError(c, err) {
	//	return
	//}
	fullPath := c.Query("path")

	err := uploadFile(fullPath, sftpClient, file)
	if err!=nil {
		common.AppError(c,"上传文件失败")
		return
	}
	common.JsonSuccess(c)

}

func uploadFile(desDir string, client *sftp.Client, header *multipart.FileHeader) error {
	if desDir == "$HOME" {
		wd, err := client.Getwd()
		if err != nil {
			return err
		}
		desDir = wd
	}
	srcFile, err := header.Open()
	if err != nil {
		return err
	}
	dstFile, err := client.Create(path.Join(desDir, header.Filename))
	if err != nil {
		return err
	}
	defer srcFile.Close()
	defer dstFile.Close()

	_, err = dstFile.ReadFrom(srcFile)
	if err != nil {
		return err
	}
	return nil
}
