package jwt

import (
	"bluebell/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var mySecret = []byte("wenhailiu@mail.ustc.edu.cn")

type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenToken(userID int64, userName string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.Cfg.JwtExpire) * time.Hour).Unix(),
			Issuer:    "wenhailiu@bluebell.com",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySecret)
}

func ParseToken(encodedToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(encodedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
