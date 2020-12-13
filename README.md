# separation
go-Gin 前后端分离

步骤二：
将一个main.go 当中实现的 注册。进行项目分离。


separation	项目根目录	
 - main.go:		主进程文件
 - routes.go:	路由文件		
     /common		数据库操作
         database.go:
     /controller		处理提交数据
 	       UserController.go:
     /model		数据结构体
  	      model.go:
     /utils		通用工具函数
   	     utils.go:
 - go.mod			
 - go.sum





