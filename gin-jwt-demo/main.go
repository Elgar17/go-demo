package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router := gin.Default()
	router.POST("/login", GenToken)     // 生成 token
	router.POST("/verify", VerifyToken) // 验证 token
	router.Run()
}

var (
	Secret     = []byte("JshH3R39") // 密钥，可以
	ExpireTime = time.Hour * 2      // token有效期
)

// JWTClaims token 中存储的信息
type JWTClaims struct {
	jwt.StandardClaims
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

// ResData 给前端返回的数据
type ResData struct {
	code  int
	msg   string
	token string
}

// GenToken 生成 token
func GenToken(c *gin.Context) {
	name, _ := c.GetPostForm("name")

	// 创建 token 中保存的数据
	claims := JWTClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpireTime).Unix(), // 过期时间
			Issuer:    "jwt_test",                        // 签发人
		},
		1,
		name,
	}
	// 创建 token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	res, err := token.SignedString(Secret)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"code":  1,
		"msg":   "ok",
		"token": res,
	})
}

// VerifyToken 验证 token
func VerifyToken(c *gin.Context) {
	tokenStr, _ := c.GetPostForm("token")
	// 解析
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		c.JSON(200, gin.H{
			"code":  -1,
			"msg":   err.Error(),
			"token": "",
		})
	}

	// 校验token
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "ok",
			"name": claims.Username,
			"id":   claims.UserID,
		})
	}
}
