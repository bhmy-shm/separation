# separation
go-Gin 前后端分离

步骤二：
将一个main.go 当中实现的 注册。进行项目分离。

- separation		项目根目录	
  - main.go:		主进程文件
  - routes.go:	路由文件
    - /common  - database.go      数据库
    - /controller  - UserController.go      数据处理
    - /model - model.go     数据结构体
    - /utils  -  utils.go     通用函数
  - go.mod
  - go.sum








