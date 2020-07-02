// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package service

import (
    "lionadmin.org/lion/common"
    "lionadmin.org/lion/model"
    "lionadmin.org/lion/request"
)

type Project struct {
    ID                  int     `json:"id"`
    SpaceId             int     `json:"space_id"`
    Name                string  `json:"name"`
    Description         string  `json:"description"`
    NeedAudit           int     `json:"need_audit"`
    Status              int     `json:"status"`
    RepoUrl             string  `json:"repo_url"`
    RepoBranch          string  `json:"repo_branch"`
    DeployMode          int     `json:"deploy_mode"`
    OnlineCluster       []int   `json:"online_cluster"`
    DeployUser          string  `json:"deploy_user"`
    DeployPath          string  `json:"deploy_path"`
    BuildScript         string  `json:"build_script"`
    BuildHookScript     string  `json:"build_hook_script"`
    DeployHookScript    string  `json:"deploy_hook_script"`
    PreDeployCmd        string  `json:"pre_deploy_cmd"`
    AfterDeployCmd      string  `json:"after_deploy_cmd"`
    AuditNotice         string  `json:"audit_notice"`
    DeployNotice        string  `json:"deploy_notice"`
    Ctime               int     `json:"ctime"`
}


type ProjectBuildScriptBind struct {
    ID                  int     `form:"id" binding:"required"`
    BuildScript         string  `form:"build_script" binding:"required"`
}

type ProjectHookScriptBind struct {
    ID                      int     `form:"id" binding:"required"`
    BuildHookScript         string  `form:"build_hook_script"`
    DeployHookScript         string  `form:"deploy_hook_script"`
}

type QueryBind struct {
    SpaceId     int     `form:"space_id"`
    Keyword     string  `form:"keyword"`
    Offset      int     `form:"offset"`
    Limit       int     `form:"limit" binding:"required,gte=1,lte=999"`
}



//func ProjectBuildScript(c *gin.Context) {
//    var form ProjectBuildScriptBind
//    if err := c.ShouldBind(&form); err != nil {
//        common.ParamError(c, err.Error())
//        return
//    }
//
//    p := &project.Project{
//        ID: form.ID,
//    }
//    if err := p.Detail(); err != nil {
//        render.AppError(c, err.Error())
//        return
//    }
//
//    if !common.InSpaceCheck(c, p.SpaceId) {
//        return
//    }
//
//    proj := &project.Project{
//        ID: form.ID,
//        BuildScript: form.BuildScript,
//    }
//    if err := proj.UpdateBuildScript(); err != nil {
//        render.AppError(c, err.Error())
//        return
//    }
//    render.Success(c)
//}

//func ProjectHookScript(c *gin.Context) {
//    var form ProjectHookScriptBind
//    if err := c.ShouldBind(&form); err != nil {
//        render.ParamError(c, err.Error())
//        return
//    }
//
//    p := &project.Project{
//        ID: form.ID,
//    }
//    if err := p.Detail(); err != nil {
//        render.AppError(c, err.Error())
//        return
//    }
//
//    if !common.InSpaceCheck(c, p.SpaceId) {
//        return
//    }
//
//    proj := &project.Project{
//        ID: form.ID,
//        BuildHookScript: form.BuildHookScript,
//        DeployHookScript: form.DeployHookScript,
//    }
//    if err := proj.UpdateHookScript(); err != nil {
//        render.AppError(c, err.Error())
//        return
//    }
//    render.Success(c)
//}

//func ProjectDelete(c *gin.Context) {
//    id := gostring.Str2Int(c.PostForm("id"))
//    if id == 0 {
//        render.ParamError(c, "id cannot be empty")
//        return
//    }
//    proj := &project.Project{
//        ID: id,
//    }
//
//    if err := proj.Detail(); err != nil {
//        render.AppError(c, err.Error())
//        return
//    }
//
//    if !common.InSpaceCheck(c, proj.SpaceId) {
//        return
//    }
//
//    if err := proj.Delete(); err != nil {
//        render.AppError(c, err.Error())
//        return
//    }
//    render.Success(c)
//}

//func ProjectDetail(c *gin.Context) {
//    id := gostring.Str2Int(c.Query("id"))
//    if id == 0 {
//        render.ParamError(c, "id cannot be empty")
//        return
//    }
//    proj := &project.Project{
//        ID: id,
//    }
//    if err := proj.Detail(); err != nil {
//        render.AppError(c, err.Error())
//        return
//    }
//
//    if !common.InSpaceCheck(c, proj.SpaceId) {
//        return
//    }
//
//    render.JSON(c, proj)
//}

//func ProjectSwitchStatus(c *gin.Context) {
//    id, status := gostring.Str2Int(c.PostForm("id")), gostring.Str2Int(c.PostForm("status"))
//    if id == 0 {
//        render.ParamError(c, "id cannot be empty")
//        return
//    }
//
//    p := &project.Project{
//        ID: id,
//    }
//    if err := p.Detail(); err != nil {
//        render.AppError(c, err.Error())
//        return
//    }
//
//    if !common.InSpaceCheck(c, p.SpaceId) {
//        return
//    }
//
//    if status !=0 {
//        status = 1
//    }
//    proj := &project.Project{
//        ID: id,
//        Status: status,
//    }
//    if err := proj.UpdateStatus(); err != nil {
//        render.AppError(c, err.Error())
//        return
//    }
//    render.Success(c)
//}



func ProjectList(info request.PageInfo) (err error, list interface{}, total int) {
    limit := info.PageSize
    offset := info.PageSize * (info.Page - 1)
    db := common.GetDB()
    var project []model.TRemoteProject
    err = db.Find(&project).Count(&total).Error
    err = db.Limit(limit).Offset(offset).Find(&project).Error
    return err, project, total
}

 func UpdateProject(project *model.TRemoteProject) (err error) {
        db := common.GetDB()
        db.Model(&project).Update()
        return err
 }

func UpdateProjectOne(project *model.TRemoteProject) (err error) {
    db := common.GetDB()
    db.Model(&project).Updates(model.TRemoteProject{BuildScript:project.BuildScript})
    return err
}

func GetProjectById(id int) (err error,result model.TRemoteProject){
    db := common.GetDB()
    var project model.TRemoteProject
    err = db.Where("id = ? ", id).First(&project).Error
    return err,project
}

func CreateProject(project *model.TRemoteProject) (err error) {
    db := common.GetDB()
    err = db.Create(&project).Error
    return err
}
