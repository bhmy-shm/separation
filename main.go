package main

import (
	"Gin/前后端分离/common"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//一、初始化数据库，并在执行完程序关闭数据库连接
	db := common.GetDB()

	defer db.Close()

	//二、开启Gin框架，到 routes.go 中解析web前端信息
	r := gin.Default()
	r = CollectRoute(r)

	//三、输出Gin日志(终端)
	panic(r.Run())
}
