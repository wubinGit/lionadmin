package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/sftp"
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/log"
	"lionadmin.org/lion/model"
	"lionadmin.org/lion/service"
	"lionadmin.org/lion/ssh"
	"lionadmin.org/lion/util/baseResult"
	"lionadmin.org/lion/util/gostring"
	"lionadmin.org/lion/xsftp"
	"os"
	"path"
	"strings"
)

/**
文件上传发布
 */
func DeployStart(c *gin.Context) {

	/**
	获取文件
	*/

	/**
	执行脚本部署命令
	*/
	id := c.Query("id")
	if len(id) == 0 {
		common.AppError(c, "参数错误")
		return
	}
	err, project := service.GetProjectById(gostring.Str2Int(id))
	if project.ID == 0 {
		common.AppError(c, "查询失败")
		return
	}

	err, machine := service.GetMachine(project.ManchineId)
	if err != nil {
		common.AppError(c, "参数错误")
		return
	}

	var tmp = "/tmp/syncd_data/tmp"
	var upload = "C:/Users/lenovo/Desktop/package/tar/2.tgz"
	client, _ := getSftpClient(c)

	UploadFile(tmp, client, upload, upload)

	sshclient, err := ssh.NewSshClient(machine)
	if err != nil {
		common.AppError(c, "系统错误")
		return
	}
	defer sshclient.Close()
	var cmds []string
	if project.PreDeployCmd != "" {
		cmds = deployScript(cmds, project.PreDeployCmd)
	}
	if project.AfterDeployCmd != "" {
		cmds = deployScript(cmds, project.AfterDeployCmd)
	}
	/**
	执行命令
	*/
	for _, cmd := range cmds {
		command, err := ssh.RunCommand(sshclient, cmd)
		if err != nil {
			log.Error(cmd)
			log.Error(command.Stderr)
		}
		log.Info(command.Stdout)
	}

}

/**
发布命令
*/
func deployScript(cmds []string, scripts string) (result []string) {
	split := strings.Split(scripts, "\n")
	for _, cmd := range split {
		cmds = append(cmds, cmd)
	}
	return cmds
}

func UploadFile(desDir string, client *sftp.Client, file string, script string) error {
	if desDir == "$HOME" {
		wd, err := client.Getwd()
		if err != nil {
			return err
		}
		desDir = wd
	}
	srcFile, err := os.Open(file)
	if err != nil {
		log.Error(err)

	}
	defer srcFile.Close()

	if err != nil {
		return err
	}

	var remoteFileName = path.Base(file)
	dstFile, err := client.Create(path.Join(desDir, remoteFileName))
	if err != nil {
		log.Error(err)
	}
	defer dstFile.Close()

	buf := make([]byte, 1024)

	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}

	return nil
}

func getSftpClient(c *gin.Context) (*sftp.Client, error) {
	client, _ := GetSftpClient(1)
	return client, nil

}

func GetSftpClient(machineId int) (*sftp.Client, error) {
	err, machine := service.GetMachine(machineId)
	if err != nil {
		return nil, err
	}
	client, err := xsftp.NewSftpClient(machine)
	return client, nil

}

/**
[cmd] $ /usr/bin/env ssh -o StrictHostKeyChecking=no -p 22 xlf@192.168.0.198 'mkdir -p /tmp/syncd_data; mkdir -p /home/syncd/deploy'
[cmd] $ /usr/bin/env sshpass -p 123456  xlf@192.168.0.198 'mkdir -p /tmp/syncd_data; mkdir -p /home/syncd/deploy'
[cmd] $ /usr/bin/env scp -o StrictHostKeyChecking=no -q -P 22 /tmp/syncd_data/tar/41.tgz xlf@192.168.0.198:/tmp/syncd_data/

sshpass -p 123456  scp /tmp/syncd_data/tmp/ptdSc83gAqfL7TToQXaWUcBW.sh  xlf@192.168.0.198:/tmp/syncd_data/

[cmd] $ /usr/bin/env ssh -o StrictHostKeyChecking=no -p 22 xlf@192.168.0.198 'cd /tmp/syncd_data; tar -zxf 41.tgz -C /home/syncd/deploy; rm -f 41.tgz'
*/

/**
发布
 */
func DeployCmdExe(c *gin.Context){
	id := c.Param("id")
	if len(id) == 0 {
		common.AppError(c, "参数错误")
		return
	}
	err, project := service.GetProjectById(gostring.Str2Int(id))
	if project.ID == 0 {
		common.AppError(c, "查询失败")
		return
	}

	err, machine := service.GetMachine(project.ManchineId)
	if err != nil {
		common.AppError(c, "参数错误")
		return
	}

	/**
	根据项目Id查询
	 */
	buildLog, err := service.SelectBuildLogStatus(project.ID)
	if buildLog.Status!= model.BUILD_SUCCESS{
		common.AppError(c, "项目没有构建")
		return
	}

	server:=&Server{
	 	ID: project.ID,
	 	Addr: machine.SshIp,
	 	User: machine.Name,
	 	Port: gostring.Str2Int(machine.SshPort),
	 	PreCmd: project.PreDeployCmd,
	 	PostCmd: project.AfterDeployCmd,
		PackFile:buildLog.Tar,
		DeployPath:project.DeployPath,
		DeployTmpPath:"/tmp/syncd_data", //先写死
	 }

	sshclient, err := ssh.NewSshClient(machine)

	//defer sshclient.Close()

	cmds := server.deployCmd()

	//执行命令
	ssh.DeployCmdExe(sshclient,cmds,&project, func() {

	})









}

func DepleyStatus(c *gin.Context)  {
	param := c.Param("id")
	DeployStatus, err := service.SelectDeployLogStatus(gostring.Str2Int(param))
	baseResult.Select(c,err,DeployStatus)
}


type Server struct {
	ID              int
	Addr            string
	User            string
	Port            int
	PreCmd          string
	PostCmd         string
	Key             string
	PackFile        string
	DeployTmpPath   string
	DeployPath      string
}



func (srv *Server) deployCmd() []string {
	var (
		useCustomKey, useSshPort, useScpPort string
	)
	if srv.Key != "" {
		useCustomKey = fmt.Sprintf("-i %s", srv.Key)
	}
	if srv.Port != 0 {
		useSshPort = fmt.Sprintf("-p %d", srv.Port)
		useScpPort = fmt.Sprintf(" -P %d", srv.Port)
	}
	var cmds []string
	if srv.PackFile == "" {
		cmds = append(cmds, "echo 'packfile empty' && exit 1")
	}

	cmds = append(cmds, []string{
		fmt.Sprintf(
			"/usr/bin/env ssh -o StrictHostKeyChecking=no %s %s %s@%s 'mkdir -p %s; mkdir -p %s'",
			useCustomKey,
			useSshPort,
			srv.User,
			srv.Addr,
			srv.DeployTmpPath,
			srv.DeployPath,
		),
		fmt.Sprintf(
			"/usr/bin/env scp -o StrictHostKeyChecking=no -q %s %s %s %s@%s:%s/",
			useCustomKey,
			useScpPort,
			srv.PackFile,
			srv.User,
			srv.Addr,
			srv.DeployTmpPath,
		),

	}...)
	if srv.PreCmd != "" {
		cmds = append(
			cmds,
			fmt.Sprintf(
				" ssh -o StrictHostKeyChecking=no %s %s %s@%s '%s'",
				useCustomKey,
				useSshPort,
				srv.User,
				srv.Addr,
				srv.PreCmd,
			),
		)
	}
	packFileName := path.Base(srv.PackFile)
	cmds = append(
		cmds,
		fmt.Sprintf(
			"/usr/bin/env ssh -o StrictHostKeyChecking=no %s %s %s@%s 'cd %s; tar -zxf %s -C %s; rm -f %s'",
			useCustomKey,
			useSshPort,
			srv.User,
			srv.Addr,
			srv.DeployTmpPath,
			packFileName,
			srv.DeployPath,
			packFileName,
		),
	)
	if srv.PostCmd != "" {
		cmds = append(
			cmds,
			fmt.Sprintf("ssh -o StrictHostKeyChecking=no %s %s %s@%s '%s'",
				useCustomKey,
				useSshPort,
				srv.User,
				srv.Addr,
				srv.PostCmd,
			),
		)
	}
	return cmds
}
