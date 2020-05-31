package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/tqsq2005/gin-vue/model"
	. "github.com/tqsq2005/gin-vue/pkg/setting"
	"time"
)

var jwtSecret []byte

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 根据用户生成token
func GenerateToken(user model.User) (string, error) {
	expiresAt := time.Now().Add(time.Duration(Config.App.JwtTokenTimeout) * time.Minute).Unix()
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    Config.App.Name,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

//解析token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}