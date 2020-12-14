package middleware

import (
	"Gin/前后端分离/common"
	"Gin/前后端分离/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

// token验证中间件
func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		// 1.获取 用户请求头中的 authorization（授权的头部） header信息
		tokenString := c.GetHeader("Authorization")
		fmt.Println("tokenString=", tokenString)

		// validata token format
		// 2.如果为空 / 不是以 Bearer 开头,则返回权限不足
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort() //将请求抛弃掉
			return
		}

		//因为 Bearer 占了7位，所以从7位往后截取
		tokenString = tokenString[7:]

		// 3. 解析token编码，用token进行判断，如果解析后的token失败，也返回权限不足
		token, claims, err := common.ParseToken(tokenString)
		// !token.Valid = 无效的token
		if !token.Valid {
			fmt.Println("token解析失败")
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token 验证错误"})
				c.Abort()
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token 时效已经过期"})
				c.Abort()
				return
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
				c.Abort()
				return
			}
		}

		// if err != nil || !token.Valid {
		// c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		// 	c.Abort()
		// 	return
		//}

		// 4. 验证通过后获取 claims中的 userID
		userID := claims.UserID
		DB := common.GetDB()
		var user model.User
		// 5.查询数据库该UERID的数据信息，并映射给结构体，将结构体数据反馈。
		sqlstr := fmt.Sprintf("select * from user where id = %d", userID)
		rows, err := DB.Query(sqlstr)
		if err != nil {
			fmt.Println("authmiddleware rows failed err=", err)
			return
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&user.ID, &user.Name, &user.Telephone, &user.Password)
			if err != nil {
				fmt.Println("authMiddleware select userid failed err=", err)
				return
			}
		}

		// 6. 如果用户ID==0
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort() //将请求抛弃掉
			return
		}

		// 如果用户存在，将user信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}
