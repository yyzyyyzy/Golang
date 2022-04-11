package common

import (
	"github.com/dgrijalva/jwt-go"
	"model"
	"time"
)

// 指定加密密钥
var jwtKey = []byte("卢本伟牛逼")

// Claims 是用户的状态和额外的元数据（官方字段）
// 如果需要额外记录别的信息，可以添加到这个结构体中
type Claims struct {
	UserId uint
	jwt.StandardClaims
}

//发放token给客户端
func ReleaseToken(user model.User) (string, error) {
	expirationtime := time.Now().Add(7 * 24 * time.Hour) //过期时间
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			//过期时间
			ExpiresAt: expirationtime.Unix(),
			//token发行日期
			IssuedAt: time.Now().Unix(),
			//token发行人
			Issuer: "李子康编程大魔王",
			//token的发行的主题
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//token.SignedString()内部生成签名字符串，再用于获取完整、已签名的token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
