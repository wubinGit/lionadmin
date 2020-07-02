package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB
/**
初始化数据库
 */
func InitDB() *gorm.DB {
	fmt.Println(viper.GetString("datasource.driverName"))
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("fail to connect databse,err:" + err.Error())
	}
	//sql 打印
	db.LogMode(true)
	//自动映射实体类到数据库
	//db.AutoMigrate(&model.User{})
	//数据库表名不加 s  设置为true
	db.SingularTable(true)
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
