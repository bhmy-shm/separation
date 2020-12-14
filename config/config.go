package config

// import (
// 	"fmt"

// 	"github.com/spf13/viper"
// )

// var config *viper.Viper

// //初始化配置文件函数
// func InitConfig() *viper.Viper {
// 	//获取当前工作目录
// 	workDir := "E:/Go语言/projects/src/Gin/separation/config"

// 	config = viper.New()
// 	//设置当前工作目录下要读取的配置文件名
// 	config.AddConfigPath(workDir)
// 	config.SetConfigName("application")
// 	config.SetConfigType("yml")

// 	err := config.ReadInConfig()
// 	if err != nil {
// 		fmt.Println("viper readconfig failed err=", err)
// 		return nil
// 	}
// 	return config
// }
