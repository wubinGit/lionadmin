package ssh

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"lionadmin.org/lion/log"
	"lionadmin.org/lion/model"
	"lionadmin.org/lion/service"
	"lionadmin.org/lion/util/gostring"
	"lionadmin.org/lion/util/utils"
	"os"
	"strings"
	"time"
)

/**
创建ssh连接  用户名密码连接或者秘钥连接
*/
func NewSshClient(manchine model.TMachine) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		Timeout:         time.Second * 5,
		User:            manchine.Name,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	SshPassword, _ := utils.Decrypt(manchine.SshPassword, []byte("xlf12345"))
	//if h.Type == "password" {
	config.Auth = []ssh.AuthMethod{ssh.Password(SshPassword)}
	//} else {
	//	config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(h.Key)}
	//}
	port := gostring.Str2Int(manchine.SshPort)
	addr := fmt.Sprintf("%s:%d", manchine.SshIp, port)
	c, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, err
	}
	return c, nil
}
func hostKeyCallBackFunc(host string) ssh.HostKeyCallback {
	hostPath, err := homedir.Expand("~/.ssh/known_hosts")
	if err != nil {
		log.Fatal("find known_hosts's home dir failed", err)
	}
	file, err := os.Open(hostPath)
	if err != nil {
		log.Fatal("can't find known_host file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				log.Error("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}
	if hostKey == nil {
		log.Error("no hostkey for %s,%v", host, err)
	}
	return ssh.FixedHostKey(hostKey)
}

func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		log.Fatal("find key's home dir failed", err)
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatal("ssh key file read failed", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}
func RunCommand(client *ssh.Client, command string) (taskresult *TaskResult, err error) {
	result := &TaskResult{
		Cmd: command,
	}
	session, err := client.NewSession()
	if err != nil {
		//log.Print(err)
		return
	}
	defer session.Close()

	var buf bytes.Buffer
	session.Stdout = &buf
	err = session.Run(command)
	if err != nil {
		result.Stderr = err.Error()
		//result.Stderr = string(buf.Bytes())
		return result, err
	}
	stdout := string(buf.Bytes())
	result.Stdout = stdout

	return result, err
}

func BuildTaskCallback(client *ssh.Client, cmds []string, project *model.TRemoteProject, tarPath string, funbuild func(result *Task)) (taskResult *Task) {
	task := &Task{}
	go func() {
		RunCommandBuild(client, project, tarPath, cmds, task)
		funbuild(task)
	}()
	log.Info(task.result)
	return task
}


func DeployCmdExe(client *ssh.Client, cmds []string, project *model.TRemoteProject, funDeploy func()) {
	go func() {
		RunCommandDeploy(client, project, cmds)
		funDeploy()
	}()
}

func RunCommandDeploy(client *ssh.Client, project *model.TRemoteProject, cmds []string)(err error) {
	buildLog := &model.TRemoteDeployLog{
		ApplyId:   project.ID,
		Status:    model.BUILD_ON,
		Ctime:     int(time.Now().Unix()),
	}
	service.CreateDeployLog(buildLog)
	for _, cmd := range cmds {
		command, err := RunCommand(client, cmd)
		log.Info(command.Stdout)
		if err != nil {
			log.Error(command.Stderr)
			buildLog.Content = gostring.JoinStrings(buildLog.Content, command.Stderr)
			buildLog.Status = model.BUILD_FILE
			buildLog.Ctime = int(time.Now().Unix())
			service.UpdateDeployLogStatus(buildLog)
			return err
		}
		buildLog.Content = gostring.JoinStrings(buildLog.Content, command.Stdout)
		buildLog.Status = model.BUILD_ON
		log.Info(command.Stdout)
		service.UpdateDeployLogStatus(buildLog)
	}
	return nil
}





func RunCommandBuild(client *ssh.Client, project *model.TRemoteProject, tarPath string, cmds []string, task *Task) (err error) {
	buildLog := &model.TRemoteBuildLog{
		ApplyId:   project.ID,
		Status:    model.BUILD_ON,
		Tar:       tarPath,
		StartTime: int(time.Now().Unix()),
		Ctime:     int(time.Now().Unix()),
	}
	service.CreateBuildLog(buildLog)
	message := map[string]interface{}{}
	for _, cmd := range cmds {
		command, err := RunCommand(client, cmd)
		log.Info(command.Stdout)
		if err != nil {
			log.Error(command.Stderr)
			buildLog.Errmsg = gostring.JoinStrings(buildLog.Errmsg, command.Stderr)
			buildLog.Status = model.BUILD_FILE
			buildLog.FinishTime = int(time.Now().Unix())
			service.UpdateBuildLogStatus(buildLog)
			return err
		}
		message["cmd"] = cmd
		message["stdout"] = command.Stdout
		fmt.Println(message)
		buildLog.Output = gostring.JoinStrings(buildLog.Output, command.Stdout)
		buildLog.Status = model.BUILD_ON
		log.Info(command.Stdout)
		service.UpdateBuildLogStatus(buildLog)
	}
	buildLog.FinishTime = int(time.Now().Unix())
	buildLog.Status = model.BUILD_FILE
	service.UpdateBuildLogStatus(buildLog)
	return nil

}

type TaskResult struct {
	Cmd     string `json:"cmd"`
	Stdout  string `json:"stdout"`
	Stderr  string `json:"stderr"`
	Success bool   `json:"success"`
}

type Task struct {
	Commands []string
	done     bool
	timeout  int
	termChan chan int
	err      error
	result   []*TaskResult
}
