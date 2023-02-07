package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type MyClaims struct {
	Phone                string `json:"phone"`
	jwt.RegisteredClaims        // 注意!这是jwt-go的v4版本新增的，原先是jwt.StandardClaims
}

func CreateToken(key string) string {
	//claim := jwt.MapClaims{
	//	"name": "tom",
	//	"age":  20,
	//}
	claim2 := MyClaims{
		Phone: "13133330000",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                       // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                       // 生效时间
		},
	}
	//第二个参数是Claims类型，实际是一个接口，放个自定义的结构体（存放实际需要传递的数据）就可以
	tokenPtr := jwt.NewWithClaims(jwt.SigningMethodHS256, claim2)
	token, err := tokenPtr.SignedString([]byte(key))
	if err != nil {
		panic(any("Token生成失败！！！"))
	}
	return token
}
