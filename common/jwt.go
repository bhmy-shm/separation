package common

import (
	"Gin/前后端分离/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//  创建一个  对称加密的 token 密钥，密钥内容自定义随便写
var jwtKey = []byte("a_sercret_crect")

//  创建一个 Claims定义用户信息
type Claims struct {
	UserID int
	jwt.StandardClaims
}

// 一、发放token
func ReleaseToken(user model.User) (string, error) {

	// 1. 设置 token 的过期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	// 2. 设置 token 第二段传递的数据信息
	claims := &Claims{
		UserID: user.ID, //传入用户id
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //token过期时间
			IssuedAt:  time.Now().Unix(),     //token发放时间
			Issuer:    "oceanlearn.tech",     //是谁发放的 Token
			Subject:   "user token",          //主题
		},
	}

	// 3. 生成token编码，用 HS256 对称方式加密，加密claims结构体内的数据。
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 二、响应API，从 tokenString 中 解析出 claims 并返回
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	//返回 token 密钥
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
