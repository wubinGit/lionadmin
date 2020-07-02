package main

import (
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"lionadmin.org/lion/common"
	"lionadmin.org/lion/route"
	"os"
)

func main() {
	InitConfig()
	db:=common.InitDB()
	defer db.Close()

	//r := gin.Default()
	//r = CollectRoute(r)
	r := route.NewRouter()
	server := common.InitServer()
	host:= server["host"]
	port := server["port"]
	if len(host)!=0 && len(port)!=0 {
		r.Run(host+ ":" +port)
	}
	//r.Run(host,port)
	//port := viper.GetString("server.port")
	//if port != "" {
	//	panic(r.Run(":" + port))
	//}
	//panic(r.Run()) // listen and serve on 0.0.0.0:8080


}
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
