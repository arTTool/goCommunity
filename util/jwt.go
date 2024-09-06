package util

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var key = []byte(GetMd5("goCommunityKey"))

func GetToken(id, email string) (string, error) {
	UserClaim := &UserClaims{
		Id:               id,
		Email:            email,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", errors.New("token签名失败")
	}
	return tokenString, nil
}

func AnalyzeToken(tokenString string) (*UserClaims, error) {
	userClaims := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, errors.New("解析token失败")
	}
	return userClaims, nil
}
