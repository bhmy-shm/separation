package common

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//数据库初始化全局变量
var DB *sql.DB

//数据库连接常量
const (
	USERNAME = "root"
	PASSWORD = "123.com"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "web"
)

// mysql 链接池函数
func InitDB() (DB *sql.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)

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

	return DB, err
}

//定义函数返回DB实例
func GetDB() *sql.DB {
	return DB
}
