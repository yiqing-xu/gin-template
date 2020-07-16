# gin-template

#### 介绍
基于golang web框架Gin搭建通用项目模板  
构建restful api

#### 软件架构
golang  
Gin  
gorm  
jwt token认证go-jwt  
gorilla-websocket
手动集成swagger

#### 项目目录
    ├─conf              配置文件  
    ├─docs              文档  
    ├─handlers          接口   
    ├─middlewares       中间件  
    ├─models            模型    
    ├─pkg               自定义  
    │  ├─jwt            jwt    
    │  └─util           工具    
    ├─routers           路由  
    ├─serializers       序列化  
    ├─static            静态文件  
    │  ├─css  
    │  ├─img  
    │  └─js  
    ├─templates         模板  
    │ .gitignore        git
    │ go.mod            go mod  
    │ go.sum            go mod
    │ main.go           main入口
    │ gin-template.exe  编译二进制文件
    │ README.md         

#### 启动
```shell script
go mod tidy     // 包管理
go mod vendor   // 同步包文件
go run main.go  // 入口
bee run   // 热重载
```

#### 备注
git conf文件夹下conf.ini配置文件
```ini
[mysql]
Type = mysql
Host = 127.0.01
Port = 3306
User = root
Password = root
DB = db
Charset = utf8mb4
Prefix = gin_

[jwt]
SecretKey =  abc

[project]
StaticUrlMapPath = {"assets/static/": "static/", "assets/docs/": "docs/"}
TemplateGlob = templates/**/*
```