package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	Secret     = []byte("JshH3R39")
	ExpireTime = time.Now().Add(time.Hour * 2).Unix()
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken
// @Description: 生成 token
// @param username 用户名
// @param password 密码
// @return string token
// @return error 错误
func GenerateToken(username, password string) (token string, err error) {

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: ExpireTime,
			Issuer:    "gin_blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString(Secret)

	return
}

// ParseToken
// @Description: 解析 token
// @param token
// @return *Claims
// @return error
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, err
		}
	}

	return nil, err
}
