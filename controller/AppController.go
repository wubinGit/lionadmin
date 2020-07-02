package controller

import (
	"github.com/gin-gonic/gin"
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/model"
	"lionadmin.org/lion/request"
	"lionadmin.org/lion/service"
	"lionadmin.org/lion/util/baseResult"
	"lionadmin.org/lion/util/gostring"
)



func CreateApp(c *gin.Context) {
	var app model.TApp
	_ = c.ShouldBind(&app)
	err := service.CreateApp(app)
	baseResult.Create(c,err)
}



func DeleteApp(c *gin.Context) {
	id := gostring.Str2Int(c.Param("id"))
	err := service.DeleteApp(id)
	baseResult.Delete(c,err)
}



func UpdateApp(c *gin.Context) {
	var app model.TApp
	_ = c.ShouldBind(&app)
	err := service.UpdateApp(&app)
	if  app.ID==0 {
		common.ParamError(c,"ID不能为空")
		return
	}
	baseResult.Update(c,err)
}



func FindApp(c *gin.Context) {
	id := gostring.Str2Int(c.Param("id"))
	err,reapp := service.GetApp(uint(id))
	baseResult.Select(c,err,reapp)
}



func GetAppList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetAppInfoList(pageInfo)
	baseResult.PageData(c,err,list,total)
}