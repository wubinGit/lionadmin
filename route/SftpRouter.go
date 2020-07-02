package route

import (
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/xsftp"
)

func SftpRouter(baseUrl string)  {
	r := Router.Group("/" + baseUrl)
	//查看所有文件
	r.GET(common.SFTP_FILE,xsftp.SftpLs)
	//查看文件详情
	r.GET(common.SFTP_CAT_FILE,xsftp.SftpCat)
	//更改文件名称
	r.GET(common.SFTP_RENAME_FILE,xsftp.SftpRename)
	//新建文件夹
	r.GET(common.SFTP_MKDIR_FILE,xsftp.SftpMkdir)
	//删除文件
	r.GET(common.SFTP_RM_FILE,xsftp.SftpRm)
	//下载文件 TODO
	r.GET(common.SFTP_DOWNLOAD_FILE,xsftp.SftpDl)

	//上传文件
	r.POST(common.SFTP_UPLOAD_FILE,xsftp.SftpUp)




}
