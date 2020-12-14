# separation
go-Gin 前后端分离

### 重点：采用 viper 第三方库，编写 yaml 配置文件，实现从配置文件中读取各种连接参数（mysql）

```yaml
server:  
  port: 1016
datasource: 
  driverName: mysql
  host: 127.0.0.1
  port: 3306
  database: web
  username: root
  password: 123.com
  charset: utf8mb4
  network: tcp
```



### 同时梳理了UserController，AuthMiddlerware中间件，在连接mysql进行查询时 与 DB 的调用关系。



#### 项目目录划分

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
- config
  - application.yml	配置文件

