## 开发手册

```shell
cd core
# 创建Api服务
goctl api new core
# 启动Api服务
go run core.go -f etc/core-api.yaml

# 使用api文件生成代码
goctl api go -api core.api -dir . -style go_zero
```
