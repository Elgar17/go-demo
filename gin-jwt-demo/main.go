package main

import (
	"gin-jwt-demo/util"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/generate", GenToken) // 生成 token
	router.POST("/parse", ParseToken)  // 解析 token
	router.Run()
}

// GenToken
// @Description: 生成 token
// @param c
func GenToken(c *gin.Context) {
	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")

	code := 400
	msg := "参数错误"
	data := make(map[string]interface{})

	if username != "" && password != "" {
		code = 200
		msg = "ok"
		data["token"], _ = util.GenerateToken(username, password) // 生成 Token
	}

	c.JSON(200, gin.H{
		"code":  code,
		"msg":   msg,
		"token": data,
	})
}

// ParseToken
// @Description: 解析 token
// @param c
func ParseToken(c *gin.Context) {
	token, _ := c.GetPostForm("token")

	code := 400
	msg := "参数错误"
	data := make(map[string]interface{})

	if token != "" {
		code = 200
		msg = "ok"
		if claims, err := util.ParseToken(token); err == nil {
			code = 200
			msg = "ok"
			data["username"] = claims.Username
			data["password"] = claims.Password
		} else {
			code = 20001
			msg = "解析 token 失败"
		}
	}

	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})

}
