package auth

import (
	"time"
	"base_auth/configs"
	jwt "github.com/dgrijalva/jwt-go"
)

func CreateJwt(username string) (string, error){
	claims:= jwt.MapClaims{}
	claims["authorized"]= true
	claims["username"]= username
	claims["exp"]= time.Now().Add(time.Hour*12).Unix()
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.EnvSecretKey()))
} 