package common

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

const (
    CODE_OK = 200
    CODE_ERROR = 400
    CODE_ERR_SYSTEM = 1000
    CODE_ERR_APP = 1001
    CODE_ERR_PARAM = 1002
	CODE_ERR_DATA_REPEAT = 1003
    CODE_ERR_LOGIN_FAILED = 1004
    CODE_ERR_NO_LOGIN = 1005
    CODE_ERR_NO_PRIV = 1006
    CODE_ERR_TASK_ERROR = 1007
	CODE_ERR_USER_OR_PASS_WRONG = 1008
	CODE_ERR_NO_DATA = 1009
)

func JSON(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, gin.H{
        "code": CODE_OK,
        "message": "success",
        "data": data,
    })
}

func PAGE(c *gin.Context, data interface{},total int) {
    c.JSON(http.StatusOK, gin.H{
        "code": CODE_OK,
        "message": "success",
        "data": data,
        "total":total,
    })
}

func CustomerError(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
        "code": code,
        "message": message,
    })
}

func RepeatError(c *gin.Context, message string) {
    c.JSON(http.StatusOK, gin.H{
        "code": CODE_ERR_DATA_REPEAT,
        "message": message,
    })
}

func NoDataError(c *gin.Context, message string) {
    c.JSON(http.StatusOK, gin.H{
        "code": CODE_ERR_NO_DATA,
        "message": message,
    })
}

func ParamError(c *gin.Context, message string) {
    c.JSON(http.StatusOK, gin.H{
        "code": CODE_ERR_PARAM,
        "message": message,
    })
}

func AppError(c *gin.Context, message string) {
    c.JSON(http.StatusOK, gin.H{
        "code": CODE_ERR_APP,
        "message": message,
    })
}


func DbError(c *gin.Context, message string) {
    c.JSON(http.StatusOK, gin.H{
        "code": CODE_ERROR,
        "message": message,
    })
}

func Success(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "code": CODE_OK,
        "message": "success",
    })
}

func JsonSuccess(c *gin.Context) {
    c.AbortWithStatusJSON(200, gin.H{"code":200,"ok": true, "msg": "success"})
}
