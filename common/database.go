package common

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	//数据库初始化全局变量
	DB *sql.DB
)

// mysql 链接池函数
func InitDB() (db *sql.DB, err error) {

	USERNAME := viper.GetString("datasource.username") //root
	PASSWORD := viper.GetString("datasource.password") //123.com
	NETWORK := viper.GetString("datasource.network")   //"tcp"
	HOST := viper.GetString("datasource.host")         //"127.0.0.1"
	PORT := viper.GetString("datasource.port")         //3306
	DATABASE := viper.GetString("datasource.database") //"web"
	CHARSET := viper.GetString("datasource.charset")   //"utf8mb4"

	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s&parseTime=True", USERNAME, PASSWORD, NETWORK, HOST, PORT, DATABASE, CHARSET)

	fmt.Printf("dsn=%v\n", dsn)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Open mysql failed err=", err)
		return
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("mysql dns err=", err)
		return
	}

	return db, err
}

//定义函数返回DB实例
func GetDB() *sql.DB {
	return DB
}
