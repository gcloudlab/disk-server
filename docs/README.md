# GCloud 开发手册

## 一、开发环境

操作系统：Windows 11

开发工具：VSCode、Docker desktop(Remote、MySQL、Redis)

前端：Node 14+、yarn 1.18.22、Vite 2.4.2

后端：Go 1.18+ 

## 二、前端

### 技术栈

- 框架：Vue3.2、Pinia
- UI库：Naive UI、tailwindcss、vue3-lottie、animate.css、vicons、sass
- 开发：Vite、Typescript
- 以及各种工具

### 目录结构

聪明的你一看就懂！



## 三、后端

### 调试命令

```shell
$ cd core
# 启动Api服务
$ go run core.go -f etc/core-api.yaml

# 使用api文件生成代码
$ goctl api go -api core.api -dir . -style go_zero
```

### 配置

#### 系统环境变量

> 自行Google：Windows配置环境变量、Linux配置环境变量

| 变量名           | 类型   | 备注                      |
| ---------------- | ------ | ------------------------- |
| MailPassword     | string | 邮箱授权码（注册服务）    |
| RedisPassword    | string | Redis 密码（验证码）      |
| TencentSecretKey | string | COS SecretKey（文件上传） |
| TencentSecretID  | string | COS SecretID（文件上传）  |

如何使用环境变量？代码详见：[define.go](/core/define/define.go)



#### 邮箱注册配置

why：提供邮箱发送验证码注册功能

目标：开启邮箱 SMTP 服务并获取**密钥**

示例：网易邮箱 (@163.com)

步骤：

1.注册网易邮箱账号：https://email.163.com/；

2.进入控制台点击顶部导航栏中的“**设置**”，弹出下拉菜单中点击“**POP3/SMTP/IMAP**”项；

3.开启服务 “**IMAP/SMTP服务**”，并将获取到的授权码保存在主机环境变量中即可。



#### 对象存储 COS 配置

why：存储用户文件资源

目标：注册并购买腾讯云 COS 服务，配置 SDK

步骤：

1.购买对象存储服务：https://console.cloud.tencent.com/cos

2.生成腾讯云密钥：https://console.cloud.tencent.com/cam/capi

2.COS SDK 开发文档： https://cloud.tencent.com/document/product/436/31215 

```shell
$ go get -u github.com/tencentyun/cos-go-sdk-v5
```



#### Redis(docker desktop)

why：提供邮箱验证码缓存功能

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

并将 `redisPassword` 设置为环境变量。



## 四、其他

### 关于本项目中为何会包含 `package.json` 文件

- 没错我用了 [`der-cli`](https://der-cli.vercel.app) 工具，需要 `package.json` 文件作**版本控制**；
- 没错 `der-cli` 是我写的一个脚手架工具，但是配置文件依赖的是 `package.json`, 所以把这文件当成配置文件即可(.config)
- 我是一只可怜又无助的小前端



## 五、部署

### 前端部署

将代码 push 到 GitHub 仓库的 **master** 分支上，由 vercel 托管平台自动拉取部署即可。

### 后端部署

**环境**：Linux - Centos 7（云服务器）

**工具**：宝塔、ssh (vscode)

**前提：配置好Linux环境变量!!**

```shell
vim /etc/bashrc
source /etc/bashrc
```

**打包**

- 1.在**本地**项目中使用 [Remote Container]() 启动一个linux容器，初始化安装Go环境；
- 2.执行 `cd core` 命令到对应路径
- 3.执行 `go build core` 打包，成功打包后生成文件名为core的二进制文件
- 4.通过宝塔面板将 core 二进制文件上传到服务器 /www/wwwroot/gcloud.aoau.top/core 目录中。（gcloud.aoau.top为nginx代理服务器地址）
- 5.执行以下操作

#### 启动Api服务

```shell
# 切换运行目录
cd /www/wwwroot/gcloud.aoau.top/core
# 在后台启动
nohup ./core &
# or 调试
./core

# 查看后台进程
ps aux|grep core
# 结束进程
kill [pid]
```
## 六、参考文档

[1]: https://golang.org/	"Go语言官网"
[2]: https://go-zero.dev/docs/quick-start/monolithic-service	"Go-Zero 单体服务"
[3]: https://gorm.io/docs	"gorm"
[4]: https://console.cloud.tencent.com/cos	"COS控制台"
[5]: https://cloud.tencent.com/document/product/436/31215	"COS开发文档"
[6]: https://console.cloud.tencent.com/cam/capi	"腾讯云密钥"
