package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//统一返回格式
// {
// 	code :20001,
// 	data:xxx,
// 	msg:xx
// }

/*
定义一个 response函数，专门用来处理数据验证后的 JSON 返回页面，
C :Gin上下文
httpstatus: http状态码
code ：自定义业务code
data : 返回数据
msg : 返回字符串
*/
func Response(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {

	// JSON 返回页面 data
	c.JSON(httpStatus, gin.H{
		"code": "code",
		"data": data,
		"msg":  msg,
	})

}

// 如果响应成果返回数据
func Success(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 200, data, msg)
}

// 如果响应失败返回数据
func Faild(c *gin.Context, msg string, data gin.H) {
	Response(c, http.StatusOK, 400, data, msg)
}
