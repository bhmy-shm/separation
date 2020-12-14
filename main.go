package main

import (
	"Gin/前后端分离/common"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {

	//初始化配置文件
	InitConfig()

	//一、初始化数据库，并在执行完程序关闭数据库连接
	db, _ := common.InitDB()

	defer db.Close()

	//二、开启Gin框架，到 routes.go 中解析web前端信息
	r := gin.Default()
	r = CollectRoute(r)

	//修改默认监听端口
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}

	//三、输出Gin日志(终端)
	panic(r.Run())
}

// //初始化配置文件函数
func InitConfig() {
	//获取当前工作目录
	workDir := "E:/Go语言/projects/src/Gin/separation/config"

	//设置当前工作目录下要读取的配置文件名
	viper.AddConfigPath(workDir)
	viper.SetConfigName("application")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("viper readconfig failed err=", err)
		return
	} else {
		fmt.Println("viper init successful")
	}
}
