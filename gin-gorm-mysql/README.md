# Gin + GORM 操作 MYSQL 数据库


## 介绍

- 连接数据库
- 数据库的增删改查

## 初始化

```bash
mkdir gin-gorm-mysql && cd
go mod init gin-gorm-mysql

# 安装 gin 框架
go get -u github.com/gin-gonic/gin
# 安装MySQL驱动
go get -u gorm.io/driver/mysql
# 安装gorm包
go get -u gorm.io/gorm
```