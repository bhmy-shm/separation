package controller

// 处理注册信息

import (
	"Gin/前后端分离/common"
	"Gin/前后端分离/dto"
	"Gin/前后端分离/model"
	"Gin/前后端分离/response"
	"Gin/前后端分离/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var (
	//注册数据实例
	user model.User
	//数据库实例
	// db = common.GetDB()
)

//验证注册信息
func Register(c *gin.Context) {

	db := common.GetDB()

	// 1.获取注册参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	// 2. 注册数据验证
	// 验证电话不能小于11位
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	//  2-1.验证密码不能小于6位
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
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
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}

	/*
		注意：密码需要加密处理，并在创建数据库结构体时将加密后的密码写入数据库。
	*/
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}

	//4. 如果用户不存在则创建该用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasePassword),
	}
	sqlstr := fmt.Sprintf("insert into user (name,telephone,password)values('%s','%s','%s');", newUser.Name, newUser.Telephone, newUser.Password)
	fmt.Println(sqlstr)
	_, err = db.Exec(sqlstr)
	if err != nil {
		fmt.Println("db insert newUser failed err=", err)
		return
	}

	// 5. 返回结果
	response.Success(c, nil, "注册成功")
}

// MySQL查询 isTelephoneExist 手机号是否存在
func IsTelephoneExist(db *sql.DB, telephone string) bool {

	//查询手机号是否存在，手机号都是唯一的。
	//如果手机号能从数据库中查出来，代表已经存在该用户
	sqlstr := `SELECT id,name,telephone,password FROM user WHERE telephone=?`
	row := db.QueryRow(sqlstr, telephone)

	//如果存在该用户，有数据映射结构体.
	row.Scan(&user.ID, &user.Name, &user.Telephone, &user.Password)

	fmt.Println("isTelephone=", user)

	//查询注册表当中 user.ID 是否 = 0，如果 ==0代表有用户信息，!=0代表没有用户信息
	if user.ID != 0 {
		return true
	}

	//用户不存在返回
	return false
}

//验证登录信息
func Login(c *gin.Context) {

	db := common.GetDB()
	// 1. 获取参数
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

	//  2-1.验证密码不能小于6位
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码必须大于6位数",
		})
		return
	}

	// 3. 判断手机号是否存在
	// 如果手机号存在就登录，不存在就反馈错误信息
	if IsTelephoneExist(db, telephone) != true {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": fmt.Sprintf("该用户：%s不存在", name)})
	}

	// 4. 判断密码是否正确	（注意用户密码不能明文保存）
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	// 5. 如果都没有问题就发送 token给前端
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error:%v", err)
	}
	// 6. 反馈结果
	response.Success(c, gin.H{"token": token}, "登录")
}

// 用户信息
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}
