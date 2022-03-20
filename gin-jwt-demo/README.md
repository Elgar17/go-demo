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
- 路径：`http://localhost:8080/generate`
- 参数：username,password
- 返回：token

返回实例：
```json
{
    "code": 200,
    "msg": "ok",
    "token": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiIxMjM0NTYiLCJleHAiOjE2NDc3NjIzMjcsImlzcyI6Imdpbl9ibG9nIn0.HEFQ_H-7SwRIzWWP6kCMjFlGjGLlSvOxdGLEjJcl4gE"
    }
}
```

（2）验证 token

- 请求方式：POST
- 路径：`http://localhost:8080/parse`
- 参数：token
- 返回：name,userid

返回实例：
```json
{
    "code": 200,
    "msg": "ok",
    "data": {
        "password": "123456",
        "username": "admin"
    }
}
```