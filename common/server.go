package common

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitServer() (server map[string]string) {
	host := viper.GetString("host")
	port := viper.GetString("server.port")
	runmode := viper.GetString("RunMode")
	gin.SetMode(runmode)
	serverMap:=map[string]string{"host":host,"port":port}
	return serverMap
}
