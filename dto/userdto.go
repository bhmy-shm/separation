package dto

import "Gin/前后端分离/model"

//	统一请求封装
// 	token验证成功后返回的 数据信息
//	只返回给前端用户名称和手机号

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

//
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
