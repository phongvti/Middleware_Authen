package auth

import (
	"base_auth/configs"
	"base_auth/models"
	"time"

	"github.com/kataras/iris/v12/middleware/jwt"
)

func CreateJwt(username string) (string, error) {
	signer := jwt.NewSigner(jwt.HS256, []byte(configs.EnvSignKey()), 60*time.Minute)
	payload := models.User{Username: username}
	token, err := signer.Sign(payload)
	return string(token), err
}
