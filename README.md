# separation
go-Gin 前后端分离

步骤二：
将一个main.go 当中实现的 注册。进行项目分离。

- separation		项目根目录	
  - main.go:		主进程文件
  - routes.go:	路由文件
    - /common  - database.go
    - /controller  - UserController.go
    - /model - model.go
    - /utils  -  utils.go
  - go.mod
  - go.sum








