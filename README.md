# Go_lib

## 项目介绍
Go_lib是一个用于管理图书馆等地的图书信息的应用程序，它可以帮助管理员更好地管理图书、借还书等功能。本图书管理系统采用了Go语言开发，主要用来练手，熟悉后端基本流程使用。

## 初始化

### 环境
```markdown
go 1.20
mysql 8.0.32
```



### 前端

```markdown
拉取代码 main分支

git clone -b main https://github.com/JiangH156/go_lib.git

更新依赖

npm install --dependencies

启动前台

yarn serve
```

### 后端
```markdown
拉取代码

git clone -b master https://github.com/JiangH156/go_lib.git

更新依赖

go mod tidy

初始化数据库

启动

go run main.go
更改config/application_dev.yml配置文件
```

## 组织结构
```markdown
go_lib
├─common          -- 通用的代码或功能
├─config          -- 配置文件
├─controller      -- 控制器（Controller）层，接受请求并处理响应
├─dto             -- 数据传输（DTO）层，处理实体传输和响应数据的转换
├─middleware      -- 中间件层，处理请求处理前后的逻辑
├─model           -- 数据库实体（Model）层，处理数据库相关操作
├─repository      -- 数据访问（Repository）层，处理数据的持久化和访问
├─response        -- 响应处理层，格式化响应数据和处理异常
├─router          -- 路由（Router）层，负责处理路由和中间件的注册
├─service         -- 服务（Service）层，处理业务逻辑
├─utils           -- 工具类和通用函数
├─vo              -- 视图对象（VO）层，封装视图数据传输格式和管理响应数据格式
└─main.go         -- 后台启动入口
```

## 技术选型
|  技术  |           功能            |                     官网                     |
| :----: | :-----------------------: | :------------------------------------------: |
|  Gin   |    Web 框架，路由注册     |         https://gin-gonic.com/zh-cn/         |
|  Gorm  | ORM框架，用于对象关系映射 |            https://gorm.io/zh_CN/            |
| MySQL  |       关系型数据库        |            https://www.mysql.com/            |
|  JWT   | 跨域认证，生成和验证令牌  |               https://jwt.io/                |
| Viper  |         配置文件          |        https://github.com/spf13/viper        |
| Gomail |         邮件服务          |     https://github.com/go-gomail/gomail      |
| Bcrypt |       密码加密服务        | https://godoc.org/golang.org/x/crypto/bcrypt |


## 功能
- 管理员
    - 删除用户
    - 添加图书
    - 修改图书信息
    - 删除借阅信息
    - 管理举报信息
    - 删除图书
    - 邮件功能
- 用户
    - 注册账号
    - 登录账号
    - 搜索图书
    - 借阅图书
    - 归还图书
    - 举报评论
    - 点赞评论

