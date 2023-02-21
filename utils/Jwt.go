package utils

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"strings"
	"time"
)

// MyClaims 用于在负荷存放用户信息
type MyClaims struct {
	UserId               string `json:"user_id"`
	UserName             string `json:"username"`
	jwt.RegisteredClaims        // 注意!这是jwt-go的v4版本新增的，原先是jwt.StandardClaims
}

// Secret
var key = []byte("三体舰队")

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return key, nil // 这是我的secret
	}
}

func CreateToken(userId int64, userName string) string {
	claim := MyClaims{
		UserId:   strconv.FormatInt(userId, 10),
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                       // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                       // 生效时间
		},
	}
	//第二个参数是Claims类型，实际是一个接口，放个自定义的结构体（存放实际需要传递的数据）就可以
	tokenPtr := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := tokenPtr.SignedString(key)
	if err != nil {
		panic(any("Token生成失败！！！"))
	}
	return token
}

// ParseToken 解析并验证Token
func ParseToken(token string) (MyClaims, bool) {
	myClaims := MyClaims{}
	claims, err := jwt.ParseWithClaims(token, &myClaims, Secret())
	if err != nil {
		fmt.Errorf("Token解析失败:%v", err)
	}
	if !claims.Valid {
		return MyClaims{}, false
	}
	return myClaims, true
}

// ValidateToken Token验证
func ValidateToken(token string) bool {
	fmt.Println("令牌为：", token)
	// 解析token，成功则返回MyClaims
	myClaims, valid := ParseToken(token)
	if !valid {
		//ctx.JSON(http.StatusUnauthorized, pb.DouyinUserResponse{
		//	StatusCode: 1,
		//	StatusMsg:  "Token解析错误，验证失败",
		//})
		return false
	}

	// 从redis获取token信息
	result, err := RedisDb.Get(context.Background(), strings.Join([]string{"token", token}, ":")).Result()
	if err != nil {
		panic(any(fmt.Sprintf("获取Redis中的Token失败：%v", err)))
	}
	tokenInfo := JsonToStruct(result)
	//redis校验token信息
	if tokenInfo.UserId == myClaims.UserId && tokenInfo.UserName == myClaims.UserName {
		return true
	}
	return false
}
