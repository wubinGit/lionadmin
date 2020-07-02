package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/sftp"
	"io/ioutil"
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/log"
	"lionadmin.org/lion/model"
	"lionadmin.org/lion/remote"
	"lionadmin.org/lion/request"
	"lionadmin.org/lion/service"
	"lionadmin.org/lion/ssh"
	"lionadmin.org/lion/util/baseResult"
	"lionadmin.org/lion/util/gofile"
	"lionadmin.org/lion/util/gostring"
	"os"
	"os/exec"
	"path"
	"strings"
)

/*
	1.拉分支
	2.生成脚本 上传脚本
	3.执行脚本 打包
	4.复制文件到远程机器
	5.执行脚本启动服务 远程执行脚本启动服务 jar
/**
rm -fr /tmp/syncd_data/tmp/9
/usr/bin/env git clone -q git@github.com:wubinGit/xlf-syncd.git /tmp/syncd_data/tmp/9
cd /tmp/syncd_data/tmp/9 && /usr/bin/env git checkout -q master
echo "Now is" `date`
echo "Run user is" `whoami`
rm -f /tmp/syncd_data/tar/9.tgz
/bin/bash -c /tmp/syncd_data/tmp/a5sdQLRXShBIgloq7SzsDVeJ.sh  //打包 脚本
rm -f /tmp/syncd_data/tmp/a5sdQLRXShBIgloq7SzsDVeJ.sh
rm -fr /tmp/syncd_data/tmp/9
echo "Compile completed" `date`

*/

/**
[cmd] $ /usr/bin/env ssh -o StrictHostKeyChecking=no -p 22 xlf@192.168.0.198 'mkdir -p /tmp/syncd_data; mkdir -p /home/syncd/deploy'
[cmd] $ /usr/bin/env scp -o StrictHostKeyChecking=no -q -P 22 /tmp/syncd_data/tar/41.tgz xlf@192.168.0.198:/tmp/syncd_data/
[cmd] $ /usr/bin/env ssh -o StrictHostKeyChecking=no -p 22 xlf@192.168.0.198 'cd /tmp/syncd_data; tar -zxf 41.tgz -C /home/syncd/deploy; rm -f 41.tgz'
*/

type BuildRequest struct {
	Id int `form:"id" json:"id"`
}

func GetProjiectList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}
	err, list, total := service.ProjectList(pageInfo)
	baseResult.PageData(c, err, list, total)
}

/**
?machineId=1&id=2
*/
func BuildProject(c *gin.Context) {
	var buildRequest BuildRequest
	err := c.ShouldBind(&buildRequest)
	if err != nil {
		common.AppError(c, "参数错误")
		return
	}
	err, project := service.GetProjectById(buildRequest.Id)
	if project.ID == 0 {
		common.AppError(c, "查询失败")
		return
	}
	_, machine := service.GetMachine(project.ManchineId)
	if machine.ID == 0 {
		common.AppError(c, "构建失败")
		return
	}

	/**
	获取项目中的 git 地址
	*/
	gitPath := project.RepoUrl

	var tmp = "/tmp/syncd_data/tmp"
	var tar = "/tmp/syncd_data/tar"

	workSpace := gostring.JoinStrings(tmp, "/", gostring.Int2Str(project.ID))
	packFile := gostring.JoinStrings(tar, "/",  gostring.Int2Str(project.ID), ".tgz")

	/**
	获取项目中的 git 地址
	*/
	repo := remote.NewRepo(gitPath, workSpace)
	branch := project.RepoBranch
	if len(branch) != 0 {
		repo.SetBranch(branch)
	} else {
		repo.SetBranch("master")
	}
	build, err := remote.NewBuild(repo, workSpace, "/tmp/syncd_data/tmp", packFile, project.BuildScript)
	if err != nil {
		common.AppError(c, err.Error())
		return
	}

	file, script := remote.CreateScriptFileShell(project.BuildScript, build)

	if err := gofile.CreateFile(script, []byte(file), 0777); err != nil {
		common.AppError(c, err.Error())
		log.Error(err.Error())
		return
	}

	client, err := GetSftpClient(project.ManchineId)
	if err != nil {
		common.AppError(c, "系统错误")
		return
	}
	err = UploadScript(script, client, script, file)
	if err != nil {
		common.AppError(c, "文件上传失败")
		return
	}
	defer client.Close()

	//构建cmd任务
	cmds := remote.BuildTask(build)

	sshclient, err := ssh.NewSshClient(machine)

	/**
	新增构建记录
	*/
	//buildLog:=&model.TRemoteBuildLog{
	//	ApplyId:project.ID,
	//	Status:model.BUILD_ON,
	//	Tar: packFile,
	//	StartTime: int(time.Now().Unix()),
	//	Ctime: int(time.Now().Unix()),
	//}
	//
	//message:= map[string]interface{}{}

	//执行远程cmd
	//for _, cmd := range cmds {
	//	//go ssh.RunCommand(sshclient, cmd);
	//	command, err := ssh.RunCommand(sshclient, cmd)
	//	if err != nil {
	//		log.Error(cmd)
	//		log.Error(command.Stderr)
	//		buildLog.Errmsg=gostring.JoinStrings(buildLog.Errmsg,command.Stderr)
	//		buildLog.Status=model.BUILD_FILE
	//		buildLog.FinishTime =int(time.Now().Unix())
	//		service.CreateBuildLog(buildLog)
	//		common.JSON(c, nil)
	//	}
	//	message["cmd"]=cmd
	//	message["errmsg"]=command.Stdout
	//	fmt.Println(message)
	//	buildLog.Output=gostring.JoinStrings(buildLog.Output,command.Stdout)
	//	buildLog.Status=model.BUILD_ON
	//	log.Info(command.Stdout)
	//}
	//buildLog.FinishTime =int(time.Now().Unix())
	//buildLog.Status=model.BUILD_FILE
	//
	//service.CreateBuildLog(buildLog)
	ssh.BuildTaskCallback(sshclient, cmds,&project,packFile, func(result *ssh.Task) {
		output := string(gostring.JsonEncode(result))
		log.Info(output)
	})

	//defer sshclient.Close()

	common.JSON(c, nil)

}

//win package
func WinBuild(c *gin.Context) {

	var buildRequest BuildRequest
	err := c.ShouldBind(&buildRequest)
	if err != nil {
		common.AppError(c, "参数错误")
		return
	}
	err, project := service.GetProjectById(buildRequest.Id)
	if project.ID == 0 {
		common.AppError(c, "查询失败")
		return
	}
	_, machine := service.GetMachine(project.ManchineId)
	if machine.ID == 0 {
		common.AppError(c, "构建失败")
		return
	}

	/**
	获取项目中的 git 地址
	*/
	gitPath := project.RepoUrl

	var tmp = "C:/Users/lenovo/Desktop/package/tmp"
	var tar = "C:/Users/lenovo/Desktop/package/tar"

	workSpace := gostring.JoinStrings(tmp, "/", gostring.Int2Str(project.ID))
	packFile := gostring.JoinStrings(tar, "/", gostring.Int2Str(project.ID), ".tgz")

	/**
	获取项目中的 git 地址
	*/
	repo := remote.NewRepo(gitPath, workSpace)
	branch := project.RepoBranch
	if len(branch) != 0 {
		repo.SetBranch(branch)
	} else {
		repo.SetBranch("master")
	}
	build, err := remote.NewBuild(repo, workSpace, "C:/Users/lenovo/Desktop/package", packFile, project.BuildScript)
	if err != nil {
		common.AppError(c, err.Error())
		return
	}

	file, script := remote.CreateScriptFile(project.BuildScript, build)
	//生成脚本文件
	if err := gofile.CreateFile(script, []byte(file), 0777); err != nil {
		common.AppError(c, err.Error())
		log.Error(err.Error())
		return
	}
	//构建cmd
	cmds := remote.BuildTaskWin(build)
	/**
	git clone  https://github.com/wubinGit/xlf-syncd.git C:/Users/lenovo/Desktop/package/tmp/2
	cd C:/Users/lenovo/Desktop/package/tmp/2
	git checkout -q master
	cd C:/Users/lenovo/Desktop/package
	C:/Users/lenovo/Desktop/package/tmp/O0BKp6sCbzJ9WROqFU8Ln0z0.sh
	cd C:/Users/lenovo/Desktop/package/tmp/2/target
	tar cvf C:/Users/lenovo/Desktop/package/tar/tmp.tgz *
	*/
	//java 构建
	for _, cmd := range cmds {
		split := strings.Split(cmd, " ")
		if len(split) == 1 {
			command := exec.Command(split[0], "")
			log.Info(command)
			_ = command.Run()
			continue
		} else {
			first := split[0]
			if strings.Compare(first, "git") == 0 {
				command := exec.Command(split[0], split[1:]...)
				log.Info(command)
				pipe, err := command.StdoutPipe()
				if err != nil {
					log.Error(err.Error())
					continue
				}
				defer pipe.Close()
				if output, err := ioutil.ReadAll(pipe); err != nil {
					log.Info(cmd)

					log.Info(output)
				}
				_ = command.Run()
				continue
			} else if strings.Compare(first, "cd") == 0 {
				command := exec.Command(split[0], split[1:]...)
				log.Info(command)
				command.Start()
				continue
			} else if strings.Compare(first, "tar") == 0 {
				command := exec.Command(split[0], split[1:]...)
				log.Info(command)
				command.Run()
				continue
			}
		}
	}

	//fmt.Println(err)
	//fmt.Println(err)

	//构建cmd任务
	//cmdtest7 := exec.Command("git", "clone","https://github.com/wubinGit/xlf-syncd.git","C:/Users/lenovo/Desktop/package/tmp/2")
	//
	////_ = cmdtest7.Start()
	//
	//_ = cmdtest7.Run()
	//log.Info(cmdtest7)
	//
	//
	//cmdtest := exec.Command("cd", "C:/Users/lenovo/Desktop/package/tmp/2")
	//_ = cmdtest.Start()
	//
	//log.Info(cmdtest)
	//cmdtest0 := exec.Command("git", "checkout","-q","master")
	//_ = cmdtest0.Start()
	//log.Info(cmdtest0)

	//cmdtest3 := exec.Command("cd ","C:/Users/lenovo/Desktop/package")
	//_ = cmdtest3.Start()
	//
	//
	//cmdtest4 := exec.Command("C:/Users/lenovo/Desktop/package/tmp/yqVnQeDPvfcpkN7JK0zvtty4.bat","")
	//_ = cmdtest4.Run()

	//log.Info(cmdtest4)
	///**
	//cd C:\Users\lenovo\Desktop\package\tmp\2\target
	//tar cvf C:\Users\lenovo\Desktop\package\tar\tmp.tgz  *
	//*/
	//cmdtest5 := exec.Command("echo \"Compile completed\" `date`","")
	//_ = cmdtest5.Start()
	//
	////
	//cmdtest0 := exec.Command("cd","C:/Users/lenovo/Desktop/package/tmp/2/target")
	//err = cmdtest0.Start()
	//
	//cmdtest := exec.Command("dir","")
	//err = cmdtest.Start()

	//
	//cmdtest8 := exec.Command("tar","cvf","C:/Users/lenovo/Desktop/package/tar/tmp.tgz","*")
	//_ = cmdtest8.Run()
	//log.Info(cmdtest8)

	//	stdout, err := cmdtest.StdoutPipe()
	//	if err != nil {
	//		log.Fatal(err)
	//
	//	}
	//	defer stdout.Close()
	//log.Info(cmdtest)
	//
	//if opBytes, err := ioutil.ReadAll(stdout); err != nil {  // 读取输出结果
	//
	//		log.Fatal(err)
	//
	//	} else {
	//
	//		log.Info(string(opBytes))
	//
	//	}

	//	for _, cmd := range cmds {
	//		log.Info(cmd)
	//		cmd := exec.Command(cmd, "")
	//		stdout, err := cmd.StdoutPipe()
	//		if err != nil {
	//			log.Fatal(err)
	//
	//		}
	//		defer stdout.Close()
	////git clone  https://github.com/wubinGit/xlf-syncd.git C:/Users/lenovo/Desktop/package/tmp/2
	//		serr := cmd.Start()
	//		if serr != nil {
	//			log.Fatal(err)
	//		}
	//		log.Info("Waiting for command to finish...")
	//		err = cmd.Wait()
	//		log.Info("Command finished with error: %v", err)
	//
	//		if opBytes, err := ioutil.ReadAll(stdout); err != nil {  // 读取输出结果
	//
	//			log.Fatal(err)
	//
	//		} else {
	//
	//			log.Info(string(opBytes))
	//
	//		}
	//	}
	//

}

func UploadScript(desDir string, client *sftp.Client, file string, script string) error {
	if desDir == "$HOME" {
		wd, err := client.Getwd()
		if err != nil {
			return err
		}
		desDir = wd
	}
	//srcFile, err := header.Open()
	srcFile, err := os.Open(file)
	if err != nil {
		log.Error(err)

	}
	defer srcFile.Close()

	if err != nil {
		return err
	}

	var remoteFileName = path.Base(file)
	dstFile, err := client.Create(path.Join("/tmp/syncd_data/tmp", remoteFileName))
	if err != nil {
		log.Error(err)
	}
	defer dstFile.Close()

	//buf := make([]byte, 1024)

	buf := []byte(script)

	//写入脚本文件
	dstFile.Write(buf)

	return nil
}

func BuildStatus(c *gin.Context)  {
	projectId := c.Param("id")
	if len(projectId)==0 {
		common.AppError(c,"参数错误")
		return
	}
	status, err := service.SelectBuildLogStatus(gostring.Str2Int(projectId))
	baseResult.Select(c,err,status)
}



func SshRuncommend(c *gin.Context) {
	var manchine model.TMachine
	manchine.ID = 1
	manchine.Name = "xlf"
	manchine.SshPassword = "fac291c45335be3b"
	manchine.SshPort = "22"
	manchine.SshIp = "192.168.0.198"
	client, err := ssh.NewSshClient(manchine)
	if err != nil {
		//return nil, err
	}
	/**
	1.拉分支
	2.执行脚本 打包
	3.复制文件到远程机器
	4.执行脚本启动服务 远程执行脚本启动服务 jar
	*/
	/**
	rm -fr /tmp/syncd_data/tmp/9
	/usr/bin/env git clone -q git@github.com:wubinGit/xlf-syncd.git /tmp/syncd_data/tmp/9
	cd /tmp/syncd_data/tmp/9 && /usr/bin/env git checkout -q master
	echo "Now is" `date`
	echo "Run user is" `whoami`
	rm -f /tmp/syncd_data/tar/9.tgz
	/bin/bash -c /tmp/syncd_data/tmp/a5sdQLRXShBIgloq7SzsDVeJ.sh  //打包 脚本
	rm -f /tmp/syncd_data/tmp/a5sdQLRXShBIgloq7SzsDVeJ.sh
	rm -fr /tmp/syncd_data/tmp/9
	echo "Compile completed" `date`

	*/

	/**
	[cmd] $ /usr/bin/env ssh -o StrictHostKeyChecking=no -p 22 xlf@192.168.0.198 'mkdir -p /tmp/syncd_data; mkdir -p /home/syncd/deploy'
	[cmd] $ /usr/bin/env sshpass -p 123456  xlf@192.168.0.198 'mkdir -p /tmp/syncd_data; mkdir -p /home/syncd/deploy'
	[cmd] $ /usr/bin/env scp -o StrictHostKeyChecking=no -q -P 22 /tmp/syncd_data/tar/41.tgz xlf@192.168.0.198:/tmp/syncd_data/

	sshpass -p 123456  scp /tmp/syncd_data/tmp/ptdSc83gAqfL7TToQXaWUcBW.sh  xlf@192.168.0.198:/tmp/syncd_data/

	[cmd] $ /usr/bin/env ssh -o StrictHostKeyChecking=no -p 22 xlf@192.168.0.198 'cd /tmp/syncd_data; tar -zxf 41.tgz -C /home/syncd/deploy; rm -f 41.tgz'
	*/

	ssh.RunCommand(client, "rm -fr /tmp/syncd_data/tmp/45")

	ssh.RunCommand(client, "git clone git@github.com:wubinGit/xlf-syncd.git /tmp/syncd_data/tmp/45")

	ssh.RunCommand(client, "cd /tmp/syncd_data/tmp/45 && git checkout -q master")

	ssh.RunCommand(client, "echo \"Now is\" `date`")

	ssh.RunCommand(client, "echo \"Run user is\" `whoami`")

	ssh.RunCommand(client, "rm -f /tmp/syncd_data/tar/45.tgz")

	ssh.RunCommand(client, "/bin/bash -c /tmp/syncd_data/tmp/NIDHKGcCabF6vP0SXFCuDPjk.sh")

	//ssh.RunCommand(client, "rm -f /tmp/syncd_data/tmp/NIDHKGcCabF6vP0SXFCuDPjk.sh")

	//ssh.RunCommand(client, "rm -fr /tmp/syncd_data/tmp/45")

	ssh.RunCommand(client, "echo \"Compile completed\" `date`")

}

/**
after_deploy_cmd: "ls"
audit_notice: "1947832562@qq.com"
deploy_mode: 0
deploy_notice: "1947832562@qq.com"
deploy_path: "/home/xlf/app"
deploy_user: "xlf"
description: "desc"
name: "hello "
need_audit: true
online_cluster: 1
pre_deploy_cmd: ""
repo_branch: "master"
repo_url: "https://syncd.cc/docs/#/build"
space_id: 1
*/
type ProjectFormBind struct {
	ID             int    `form:"id" json:"id"`
	SpaceId        int    `form:"space_id" json:"space_id"`
	Name           string `form:"name" json:"name" binding:"required"`
	Description    string `form:"description" json:"description"`
	NeedAudit      int    `form:"need_audit" json:"need_audit"`
	RepoUrl        string `form:"repo_url" json:"repo_url" binding:"required"`
	RepoBranch     string `form:"repo_branch" json:"repo_branch"`
	DeployMode     int    `form:"deploy_mode" json:"deploy_mode" binding:"required"`
	OnlineCluster  int    `form:"online_cluster" json:"online_cluster" binding:"required"`
	DeployUser     string `form:"deploy_user" json:"deploy_user" binding:"required"`
	DeployPath     string `form:"deploy_path" json:"deploy_path" binding:"required"`
	PreDeployCmd   string `form:"pre_deploy_cmd" json:"pre_deploy_cmd"`
	AfterDeployCmd string `form:"after_deploy_cmd" json:"after_deploy_cmd"`
	AuditNotice    string `form:"audit_notice" json:"audit_notice"`
	DeployNotice   string `form:"deploy_notice" json:"deploy_notice"`
	MachineId      int    `form:"machine_id" json:"machine_id"`
}

func ProjectCreateOrUpdate(c *gin.Context) {
	var projectForm ProjectFormBind
	if err := c.ShouldBind(&projectForm); err != nil {
		common.ParamError(c, err.Error())
		return
	}

	repoBranch := projectForm.RepoBranch
	if projectForm.DeployMode == 2 {
		repoBranch = ""
	}
	fmt.Println(repoBranch)
	project := &model.TRemoteProject{
		ID:             projectForm.ID,
		Name:           projectForm.Name,
		Description:    projectForm.Description,
		NeedAudit:      projectForm.NeedAudit,
		RepoUrl:        projectForm.RepoUrl,
		DeployMode:     projectForm.DeployMode,
		RepoBranch:     projectForm.RepoBranch,
		OnlineCluster:  gostring.Int2Str(projectForm.OnlineCluster),
		DeployUser:     projectForm.DeployUser,
		DeployPath:     projectForm.DeployPath,
		PreDeployCmd:   projectForm.PreDeployCmd,
		AfterDeployCmd: projectForm.AfterDeployCmd,
		AuditNotice:    projectForm.AuditNotice,
		DeployNotice:   projectForm.DeployNotice,
		ManchineId:     projectForm.OnlineCluster,
	}

	if err := service.CreateProject(project); err != nil {
		common.AppError(c, err.Error())
		return
	}
	common.Success(c)
}
func GetProjectById(c *gin.Context) {
	param := c.Param("id")
	if len(param) == 0 {
		common.AppError(c, "参数错误")
		return
	}
	err, result := service.GetProjectById(gostring.Str2Int(param))
	baseResult.Select(c, err, result)
}

/**
build_script: "cd ${env_workspace}↵mvn -U clean install -Dmaven.test.skip=true -Dmaven.javadoc.skip=true ↵cd ${env_workspace}/target↵tar -zcvf ${env_pack_file} *"
id: 2
*/
type UpdateProjectRequest struct {
	BuildScript string `json:"build_script" form:"build_script"`
	Id          int    `json:"id" form:"id"`
}

func UpdateProject(c *gin.Context) {
	var project UpdateProjectRequest
	err := c.ShouldBind(&project)
	if err != nil {
		common.AppError(c, "参数错误")
		return
	}
	if project.Id == 0 || len(project.BuildScript) == 0 {
		common.AppError(c, "非法参数")
		return
	}
	err, result := service.GetProjectById(project.Id)
	result.BuildScript = project.BuildScript
	err = service.UpdateProjectOne(&result)
	baseResult.Update(c, err)

}

//func (p *Project) CreateOrUpdate() error {
//	project := &model.TSydProject{
//		ID: p.ID,
//		SpaceId: p.SpaceId,
//		Name: p.Name,
//		Description: p.Description,
//		NeedAudit: p.NeedAudit,
//		RepoUrl: p.RepoUrl,
//		DeployMode: p.DeployMode,
//		RepoBranch: p.RepoBranch,
//		OnlineCluster: gostring.JoinIntSlice2String(p.OnlineCluster, ","),
//		DeployUser: p.DeployUser,
//		DeployPath: p.DeployPath,
//		PreDeployCmd: p.PreDeployCmd,
//		AfterDeployCmd: p.AfterDeployCmd,
//		AuditNotice: p.AuditNotice,
//		DeployNotice: p.DeployNotice,
//	}
//	if project.ID > 0 {
//		updateData := map[string]interface{}{
//			"name": p.Name,
//			"description": p.Description,
//			"need_audit": p.NeedAudit,
//			"repo_url": p.RepoUrl,
//			"deploy_mode": p.DeployMode,
//			"repo_branch": p.RepoBranch,
//			"online_cluster": gostring.JoinIntSlice2String(p.OnlineCluster, ","),
//			"deploy_user": p.DeployUser,
//			"deploy_path": p.DeployPath,
//			"pre_deploy_cmd": p.PreDeployCmd,
//			"after_deploy_cmd": p.AfterDeployCmd,
//			"audit_notice": p.AuditNotice,
//			"deploy_notice": p.DeployNotice,
//		}
//		//if ok := project.UpdateByFields(updateData, model.QueryParam{
//		//	Where: []model.WhereParam{
//		//		model.WhereParam{
//		//			Field: "id",
//		//			Prepare: project.ID,
//		//		},
//		//	},
//		//}); !ok {
//		//	return errors.New("project update failed")
//		//}
//		fmt.Println(updateData)
//	} else {
//		if ok := project.Create(); !ok {
//			//return errors.New("project create failed")
//			 log.Error("project create failed")
//		}
//	}
//	return nil
//}
