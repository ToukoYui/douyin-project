package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"testing"
	"time"
)

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZSI6IjEzMTMzMzMwMDAwIiwiZXhwIjoxNjc1NjE0NjIxLCJuYmYiOjE2NzU2MDM4MjEsImlhdCI6MTY3NTYwMzgyMX0.fS8LMdEjuz03ER4WQmF0LPtNahUwr09lTMKMLusiAY0"
var myKey = []byte("my_key")

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("my_key"), nil // 这是我的secret
	}
}

type MyClaims struct {
	Phone                string `json:"phone"`
	jwt.RegisteredClaims        // 注意!这是jwt-go的v4版本新增的，原先是jwt.StandardClaims
}

func TestCreate(t *testing.T) {
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
	token, err := tokenPtr.SignedString(myKey)
	fmt.Println(token, err)
}

func TestParse(t *testing.T) {
	key, _ := jwt.Parse(token, Secret())

	fmt.Println(key.Claims, key.Method, key.Signature, key.Header, key.Raw)
}
