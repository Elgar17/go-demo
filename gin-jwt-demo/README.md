# Gin JWT demo

## 功能

- 生成 token
- 验证 token

## 初始化

（1）初始化项目

```bash
go mod init gin-jwt-demo
```

（2）安装框架和包

```bash
go get github.com/dgrijalva/jwt-go
go get github.com/gin-gonic/gin
```

## 启动项目

（1）终端进入 gin-jwt-demo

```bash
go mod tidy     # 安装依赖
go run main.go  # 启动项目
```

## 接口说明

（1）生成 token

- 请求方式：POST
- 路径：`/login`
- 参数：name
- 返回：token

（2）验证 token

- 请求方式：POST
- 路径：`/verify`
- 参数：token
- 返回：name,userid