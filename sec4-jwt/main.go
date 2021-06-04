package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpiredDuration = time.Hour * 2 // 2小时
var MySecret = []byte("my secret words...")

func GenToken(usernane string) (string, error) {
	c := MyClaims{
		"qiaocc",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpiredDuration).Unix(), // 过期时间
			Issuer:    "my-project",                                // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(MySecret)
}

func parseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
func main() {
	token, err := GenToken("qiaocc")
	if err != nil {
		return
	}
	fmt.Printf("gen token, token:%v\n\n", token)

	c, err := parseToken(token)
	if err != nil {
		return
	}
	fmt.Println(c.Username, c.Issuer, c.ExpiresAt)
}
