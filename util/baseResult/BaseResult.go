package baseResult

import (
	"github.com/gin-gonic/gin"
	"lionadmin.org/lion/common"
)

func Create(c *gin.Context,err error)  {
	if err != nil {
		common.AppError(c,"创建失败")
	} else {
		common.JSON(c,"创建成功");
	}
}

func Update(c *gin.Context,err error)  {
	if err != nil {
		common.AppError(c,"更新失败")
	} else {
		common.JSON(c,"更新成功");
	}
}

func Select(c *gin.Context,err error,data interface{})  {
	if err!=nil {
		common.AppError(c,"查询失败")
	}else{
		common.JSON(c,data)
	}
}
func Delete(c *gin.Context,err error)  {
	if err != nil {
		common.AppError(c,"删除失败")
	} else {
		common.JSON(c,"删除成功");
	}
}
func PageData(c *gin.Context,err error,data interface{} ,total int)  {
	if err != nil {
		common.AppError(c, "获取数据失败")
	} else {
		common.PAGE(c, data, total)
	}
}

func Auth(c *gin.Context,err error)  {
	if err != nil {
		common.AppError(c,"授权失败")
	} else {
		common.JSON(c,"授权成功");
	}
}
