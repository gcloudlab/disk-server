# GCloud 开发手册

## 开发环境

操作系统：Windows 11

开发工具：VSCode、Navicat、Postman、Docker desktop(MySQL、Redis)



## 调试命令

```shell
$ cd core
# 创建Api服务
$ goctl api new core
# 启动Api服务
$ go run core.go -f etc/core-api.yaml

# 使用api文件生成代码
$ goctl api go -api core.api -dir . -style go_zero
```



## 配置

### 环境变量

| 变量名           | 类型   | 备注          |
| ---------------- | ------ | ------------- |
| MailPassword     | string | 邮箱密钥      |
| RedisPassword    | string | Redis 密码    |
| TencentSecretKey | string | COS SecretKey |
| TencentSecretID  | string | COS SecretID  |

代码详见：[define.go](/core/define/define.go)



### 邮箱注册配置

目标：开启 SMTP 服务并获取**密钥**

示例：网易邮箱 (@163.com)



### 对象存储 COS 配置

目标：注册并购买腾讯云 COS 服务，配置 SDK

[1]: https://console.cloud.tencent.com/cam/capi	"腾讯云密钥申请"





```shell
$ go get -u github.com/tencentyun/cos-go-sdk-v5
```



### Redis(docker desktop)

```shell
# 安装并启动一个redis容器
$ docker pull redis
$ docker run --name gredis -p 6379:6379 redis --requirepass "redisPassword"

# 进入容器 (cmd)
$ docker exec -it gredis bash
# 进入容器 (bash)
> redis-cli
# 登陆
127.0.0.1:6379> auth redisPassword
# 查看版本
127.0.0.1:6379> info
```

将 `redisPassword` 设置为环境变量。



## 其他

### 关于本项目中为何会包含 `package.json` 文件

- 因为我用了 [`der-cli`](https://der-cli.vercel.app) 工具，需要 `package.json` 文件作版本控制
- 没错 `der-cli` 是我写的一个脚手架工具，但是配置文件依赖的是 `package.json`, 所以你把这文件当成配置文件即可(.config)
- 我是一只可怜又无助的小前端



## 参考文档

Go: https://golang.org/

gorm: https://gorm.io/docs

COS: https://console.cloud.tencent.com/cos

COS SDK: https://cloud.tencent.com/document/product/436/31215

腾讯云密钥: https://console.cloud.tencent.com/cam/capi
