package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/model"
	"net/http"
	"strings"
)
/**
拦截器
 */
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		wsAuth := ctx.Query("Authorization")
		request := ctx.Request
		url := request.URL
		path := url.Path
		///api/v1/admin/ssh/connect/1
		fmt.Println(path)
		//login path
		var loginPath="/"+common.BASE_URL+common.LOGIN

		//ssh path
		var wsPath="/"+common.SSH_BASE_URL+common.SSH_CONNECT

		/**
			SFTP_DOWNLOAD_FILE="/sftp/:id/dl"
		SFTP_UPLOAD_FILE="/sftp/:id/up"
		 */
		//download path
		//var downloadPath="/"+common.SFTP_BASE_URL+common.SFTP_DOWNLOAD_FILE
		////       /api/v1/admin/sftp/sftp/1/dl
		//downindex := strings.LastIndex(wsPath, "/")
		//if len(path)>downindex {
		//	downloadPath=downloadPath[downindex:]
		//}
		//
		////upload path
		//var uploadPath="/"+common.SFTP_BASE_URL+common.SFTP_UPLOAD_FILE
		//
		//uploadindex := strings.LastIndex(wsPath, "/")
		//if len(path)>uploadindex {
		//	downloadPath=downloadPath[:uploadindex]
		//}
		//suffix := strings.HasSuffix(path, "/dl")

		if tokenString=="" && strings.Compare(path,loginPath)==0 || strings.HasSuffix(path, "/dl")  || strings.HasSuffix(path, "/up"){
			ctx.Next()
			return
		}



		index := strings.LastIndex(wsPath, "/")
		wsPath=wsPath[:index]

		if  strings.Contains(path,wsPath) {
			tokenString = wsAuth[7:]
			token, claims, err := common.ParseToken(tokenString)
			if err != nil || !token.Valid {
				ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
				ctx.Abort()
				return
			}

			//验证通过后获取Claiim中的userId
			userId := claims.UserId
			DB := common.GetDB()
			var user model.TAdmin
			DB.First(&user, userId)

			//用户不存在
			if user.ID == 0 {
				ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
				ctx.Abort()
				return
			}
			//用户存在，将user信息写入上下文
			ctx.Set("user", user)
			ctx.Next()
			return

		}





		//vcalidate token formate
		if tokenString == ""  || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//验证通过后获取Claiim中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.TAdmin
		DB.First(&user, userId)

		//用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//用户存在，将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
