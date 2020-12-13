package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var user User

//
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

const (
	USERNAME = "root"
	PASSWORD = "123.com"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "web"
)

func main() {

	db, err := InitDB()
	if err != nil {
		fmt.Println("initdb mysql failed err=", err)
	}

	defer db.Close()

	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {

		// 1.获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")

		// 2. 数据验证
		// 验证电话不能小于11位
		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "手机号必须为11位",
			})
			return
		}

		//验证密码不能小于6位
		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "密码必须大于6位数",
			})
			return
		}

		//验证用户名是否为空，如果位空给一个10位的随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}

		log.Println(name, telephone, password)

		// 3. 判断手机号是否存在
		// 如果用户存在则不允许注册, 返回 false 代表不存在，执行注册用户
		if isTelephoneExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "用户已经存在",
			})
			return
		}

		//4. 如果用户不存在则创建该用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		sqlstr := fmt.Sprintf("insert into user (name,telephone,password)values('%s','%s','%s');", newUser.Name, newUser.Telephone, newUser.Password)
		fmt.Println(sqlstr)
		_, err := db.Exec(sqlstr)
		if err != nil {
			fmt.Println("db insert newUser failed err=", err)
			return
		}

		// 5. 返回结果
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})

	panic(r.Run())
}

// 返回一个随机默认字符串
func RandomString(n int) string {

	// 定义一个byte字符串，切片类型
	var letters = []byte("asdddddWDASFSDFMLKLZXdddafsdfsdfsknfdsmldklsaAFAFSDMLMSADLd")

	// 一个新的切片类型
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())

	//循环10次，每次生成一个数字，追加到 result[i]
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

// mysql 链接池
func InitDB() (db *sql.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Open mysql failed err=", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("mysql dns err=", err)
		return
	}

	return
}

// MySQL查询 isTelephoneExist 手机号是否存在
func isTelephoneExist(db *sql.DB, telephone string) bool {

	//查询手机号是否存在，手机号都是唯一的。
	//如果手机号能从数据库中查出来，代表已经存在该用户
	sqlstr := `SELECT telephone FROM user WHERE telephone=?`
	rows, err := db.Query(sqlstr, telephone)
	//如果该用户不存在，则返回false
	if err != nil {
		fmt.Println("telephone failed err=", err)
		return false
	}
	//如果存在该用户，有数据映射结构体，则反馈一个 true 代表用户已经存在
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Telephone, &user.Password)
		if err != nil {
			fmt.Println("scan failed err=", err)
			return true
		}
	}

	defer rows.Close()
	fmt.Println("isTelephone=", user)

	//查询注册表当中 user.ID 是否 = 0，如果 ==0代表有用户信息，!=0代表没有用户信息
	if user.ID != 0 {
		return true
	}

	//用户不存在返回
	return false
}
