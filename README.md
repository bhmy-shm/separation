# separation
go-Gin 前后端分离
### 实现功能

- #### 实现用户登录请求。

- #### 通过 twj 对称加密，响应客户端登录请求，如果验证通过发送 token 编码。

- #### 创建中间件，以中间件的形式进行登录验证。

- #### 封装统一的请求返回格式，JSON反馈格式



### 项目目录划分

#### separation 

- main.go	主进程
- routers.go    路由RESTful
- model
  - model.go	用户注册，登录数据结构体
- controller
  - UserController.go	处理注册，登录，数据校验，返回用户信息
- common
  - database.go	初始化数据库
  - jwt.go    *重点：jwt 对称加密认证，①HS256的token编码，②客户端发送token解密，并根据 jwt.Token 反馈用户数据。
- middleware   （中间件）
  - AuthMiddleware.go	处理用户登录请求的 token 编码验证中间件（存在问题：会将数据库所有字段都反馈）
- dto
  - userdto.go	专门处理用户登录后，反馈得数据信息，用户名，电话号
- response
  - response.go	处理所有的 http - JSON 反馈信息，封装成函数的形式。



### 技术难点

jwt 对称加密，非对称加密，以中间件的形式处理客户端请求头部中的 token 编码。

