package controller

// 处理注册信息

import (
	"Gin/前后端分离/common"
	"Gin/前后端分离/model"
	"Gin/前后端分离/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	//注册数据实例
	user model.User
	//数据库实例

	db, _ = common.InitDB()
)

//验证注册信息
func Register(c *gin.Context) {

	// // 0.初始化数据库
	// db, err := common.InitDB()
	// if err != nil {
	// 	fmt.Println("initdb mysql failed err=", err)
	// }

	// 1.获取注册参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	// 2. 注册数据验证
	// 验证电话不能小于11位
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "手机号必须为11位",
		})
		return
	}

	//  2-1.验证密码不能小于6位
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码必须大于6位数",
		})
		return
	}

	//  2-2. 验证用户名是否为空，如果位空给一个10位的随机字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}

	log.Println(name, telephone, password)

	// 3. 判断手机号是否存在
	// 如果用户存在则不允许注册, 返回 false 代表不存在，执行注册用户
	if IsTelephoneExist(db, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户已经存在",
		})
		return
	}

	//4. 如果用户不存在则创建该用户
	newUser := model.User{
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
}

// MySQL查询 isTelephoneExist 手机号是否存在
func IsTelephoneExist(db *sql.DB, telephone string) bool {

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
