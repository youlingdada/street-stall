package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte = []byte("fdhsafhdasjkhfasjk")

/*
	构造一个含有用户信息的token
*/
func GetJwtToken(uuid string) (*string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(2 * time.Hour)

	claims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer:    "youlingdada",
		Subject:   uuid,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	temp, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}
	res := &temp
	return res, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*string, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid { // 校验token
		return &claims.Subject, nil
	}
	return nil, errors.New("invalid token")
}
