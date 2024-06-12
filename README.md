# es-client
es-client

## 开发

```shell
# 生成swagger
swag init

# 启动
go run main.go
```

## 打包

```shell
# 下载前端项目 https://github.com/liusanp/es-client-web
# 构建后将dist目录复制到public目录下
# 执行构建
go build -o build/es-client.exe
```