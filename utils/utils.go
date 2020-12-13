package utils

//工具包

import (
	"math/rand"
	"time"
)

// 工具1：返回一个随机默认字符串
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
