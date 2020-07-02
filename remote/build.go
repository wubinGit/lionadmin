package remote

import (
    "fmt"
    "lionadmin.org/lion/ssh"
    "lionadmin.org/lion/util/gostring"
    "strings"
)

type Build struct {
    repo        *Repo
    local       string
    tmp         string
    packFile    string
    scriptFile  string
    result      *Result
    task        *ssh.Task
}

const (
    STATUS_INIT = 1
    STATUS_ING = 2
    STATUS_DONE = 3
    STATUS_FAILED = 4
)

const (
    COMMAND_TIMEOUT = 86400
)

func NewBuild(repo *Repo, local, tmp, packFile, scripts string) (*Build, error) {
    build := &Build{
        repo: repo,
        local: local,
        tmp: tmp,
        packFile: packFile,
        result: &Result{
            status: STATUS_INIT,
        },
    }
    //if err := build.createScriptFile(scripts); err != nil {
    //    return build, err
    //}

   // build.initBuildTask()


    return build, nil
}
/**
env_workspace=/tmp/syncd_data/tmp/42
env_pack_file=/tmp/syncd_data/tar/42.tgz


cd ${env_workspace}
mvn -U clean install -Dmaven.test.skip=true -Dmaven.javadoc.skip=true
cd ${env_workspace}/target
tar -zcvf ${env_pack_file} *

 */
/**
win
 */
func CreateScriptFile(scripts string,b *Build) (result string,script string) {
    b.scriptFile = gostring.JoinStrings("C:/Users/lenovo/Desktop/package/tmp", "/", gostring.StrRandom(24), ".bat")
    s := gostring.JoinStrings(
        fmt.Sprintf("set env_workspace=%s\n", b.local),
        fmt.Sprintf("set env_pack_file=%s\n", b.packFile),
        scripts,
    )
    return s,b.scriptFile
}


func CreateScriptFileShell(scripts string,b *Build) (result string,script string) {
    b.scriptFile = gostring.JoinStrings("C:/Users/lenovo/Desktop/tmp", "/", gostring.StrRandom(24), ".sh")
    s := gostring.JoinStrings(
        "#!/bin/bash\n\n",
        "#--------- build scripts env ---------\n",
        fmt.Sprintf("env_workspace=%s\n", b.local),
        fmt.Sprintf("env_pack_file=%s\n", b.packFile),
        scripts,
    )
    return s,b.scriptFile
}

func BuildTask(b *Build)(cmd  []string) {
    cmds := b.repo.Fetch()
    //b.scriptFile = gostring.JoinStrings("C:/Users/lenovo/Desktop/package/tmp", "/", gostring.StrRandom(24), ".sh")
   // b.scriptFile = gostring.JoinStrings("/tmp/syncd_data/tmp", "/", gostring.StrRandom(24), ".sh")
    //b.scriptFile =     b.scriptFile=gostring.StrRandom(24), ".sh")
    split := strings.Split(b.scriptFile,"/")
    s := split[len(split)-1]
    b.scriptFile=gostring.JoinStrings("/tmp/syncd_data/tmp", "/",s)
    cmds = append(cmds, []string{
        "echo \"Now is\" `date`",
        "echo \"Run user is\" `whoami`",
        fmt.Sprintf("rm -f %s", b.packFile),
        ///bin/bash -c /tmp/syncd_data/tmp/2
        fmt.Sprintf("cd %s",b.tmp),
        //window 下测试 修改文件的权限
        fmt.Sprintf("chmod 777 %s",b.scriptFile),
        fmt.Sprintf("/bin/bash -c %s", b.scriptFile),
      //  fmt.Sprintf(".%s", b.scriptFile),
        fmt.Sprintf("rm -fr %s", b.local),
        "echo \"Compile completed\" `date`",
    }...)
    return  cmds;
}





func BuildTaskWin(b *Build)(cmd  []string) {
    cmds := b.repo.FetchWin()
    //index := strings.LastIndex(b.scriptFile, "/")
    //split := strings.Split(b.scriptFile, "/")
    //var substr=split[len(split)-1]
    //  b.scriptFile=gostring.JoinStrings("C:/Users/lenovo/Desktop/package/tmp/",substr)
    cmds = append(cmds, []string{
        fmt.Sprintf("cd %s",b.tmp),
        //window 下测试 修改文件的权限
        //fmt.Sprintf("chmod 777 %s",b.scriptFile),
        // fmt.Sprintf("/bin/bash -c %s", b.scriptFile),
        fmt.Sprintf("%s", b.scriptFile),

       	//cmdtest6 := exec.Command("cd","C:/Users/lenovo/Desktop/package/tmp/2/target")
       	//cmdtest8 := exec.Command("tar","cvf","C:/Users/lenovo/Desktop/package/tar/tmp.tgz","*")

        fmt.Sprintf("cd %s/target",b.local),
        fmt.Sprintf("tar cvf %s %s/target/*", b.packFile,b.local),
        //  fmt.Sprintf("rm -fr %s", b.local),
    }...)
    return  cmds;
    //b.task = command.NewTask(cmds, COMMAND_TIMEOUT)
}







func (b *Build) Result() *Result {
    return b.result
}



func (b *Build) PackFile() string {
    return b.packFile
}


type Result struct {
    err     error
    status  int
    stime   int
    etime   int
}

func (r *Result) During() int {
    return r.etime - r.stime
}

func (r *Result) Status() int {
    return r.status
}

func (r *Result) GetError() error {
    return r.err
}