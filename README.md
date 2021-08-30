# Go Gin Example

An example of gin contains many useful features

[简体中文](https://github.com/Quons/go-gin-example/blob/master/README_ZH.md)

## Installation
```
$ go get github.com/Quons/go-gin-example
```

## How to run

### Required

- Mysql
- Redis

### Conf

```
[database]
Type = mysql
User = root
Password =
Host = 127.0.0.1:3306
Name = blog
TablePrefix = blog_

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200
...
```

### Ready

Create a **blog database** and import [SQL](https://github.com/Quons/go-gin-example/blob/master/docs/sql/blog.sql)

### Run
```
$ cd $GOPATH/src/go-gin-example

$ go run main.go 
```

Project information and existing API

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /auth                     --> github.com/Quons/go-gin-example/routers/api.GetAuth (3 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/Quons/go-gin-example/vendor/github.com/swaggo/gin-swagger.WrapHandler.func1 (3 handlers)
[GIN-debug] GET    /api/v1/tags              --> github.com/Quons/go-gin-example/routers/api/v1.GetTags (4 handlers)
[GIN-debug] POST   /api/v1/tags              --> github.com/Quons/go-gin-example/routers/api/v1.AddTag (4 handlers)
[GIN-debug] PUT    /api/v1/tags/:id          --> github.com/Quons/go-gin-example/routers/api/v1.EditTag (4 handlers)
[GIN-debug] DELETE /api/v1/tags/:id          --> github.com/Quons/go-gin-example/routers/api/v1.DeleteTag (4 handlers)
[GIN-debug] GET    /api/v1/articles          --> github.com/Quons/go-gin-example/routers/api/v1.GetArticles (4 handlers)
[GIN-debug] GET    /api/v1/articles/:id      --> github.com/Quons/go-gin-example/routers/api/v1.GetArticle (4 handlers)
[GIN-debug] POST   /api/v1/articles          --> github.com/Quons/go-gin-example/routers/api/v1.AddArticle (4 handlers)
[GIN-debug] PUT    /api/v1/articles/:id      --> github.com/Quons/go-gin-example/routers/api/v1.EditArticle (4 handlers)
[GIN-debug] DELETE /api/v1/articles/:id      --> github.com/Quons/go-gin-example/routers/api/v1.DeleteArticle (4 handlers)

Listening port is 8000
Actual pid is 4393
```
Swagger doc

![image](https://sfault-image.b0.upaiyun.com/286/780/2867807553-5aae27c4ac806_articlex)

## Features

- RESTful API
- Gorm
- Swagger
- logging
- Jwt-go
- Gin
- Graceful restart or stop (fvbock/endless)
- App configurable
- Cron
- Redis


##配置文件管理
https://github.com/go-ini/ini
配置文件分为dev ，test ，pre ，prod 四种模式，更改配置需要同时维护
启动时通过runmode 标签进行选择，如 ./main -runmode=test

##日志系统
使用最流行的golang 日志框架logrus
https://github.com/sirupsen/logrus

##缓存
缓存使用redigo

##数据库 orm
使用gorm
https://github.com/jinzhu/gorm
放弃beego orm 的理由 1，关联查询采用join语句，无法通过sql审核。2，不支持prepare，

##依赖管理
golang 官方依赖包管理工具
https://github.com/golang/dep
安装go get -u github.com/golang/dep/cmd/dep
使用 dep ensure

##golang 官方包下载方式
以go get golang.org/x/net 为例，首先将golang.org/x 替换成 github.com/golang执行go get github.com/golang/net 命令
get完成之后执行 cp -rf $GOPATH/src/github.com/golang/* $GOPATH/src/golang.org/x/
原理为先下载官方包在github上的镜像包，然后复制到golang/x目录下

##自动文档gin-swagger
https://github.com/swaggo/gin-swagger
安装swag 工具
go get -u github.com/swaggo/swag/cmd/swag
官方包安装参考上一条

##请求数据验证
gin内置验证器为go-playground
https://github.com/go-playground/validator

#todo gorm 关联查询sql语句测试
#缓存使用
#包复用测试
#grpc测试
#response封装
#vendor dep优势
#单元测试

#todo mockito