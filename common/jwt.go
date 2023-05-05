package common

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserClaims struct {
	UserId string
	jwt.RegisteredClaims
}

// JWT过期时间
const TokenExpireDuration = time.Hour * 24

// 用于签名的字符串 一般使用随机字符串
var jwtScret = []byte("Go_lib")

// GenToken
// @Description 生成Token
// @Author John 2023-05-05 11:15:35
// @Param userId
// @Return string
// @Return error
func GenToken(userId string) (string, error) {
	//创建一个自己的声明
	claims := UserClaims{
		userId,
		jwt.RegisteredClaims{
			//发行者
			Issuer: "HuangJ",
			//主题
			Subject: "Go_lib_TOKEN",
			//过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			//生效时间
			NotBefore: jwt.NewNumericDate(time.Now()),
			//发行时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			//唯一标识符
			//ID:
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token)
	//使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, err := token.SignedString(jwtScret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken
// @Description 解析Token
// @Author John 2023-05-05 11:15:56
// @Param tokenString
// @Return *UserClaims
// @Return error
func ParseToken(tokenString string) (*UserClaims, error) {
	// 自定义claims使用
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtScret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("权限不够")
}
