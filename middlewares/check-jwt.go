package middlewares

import (
	"github.com/kataras/iris/v12"
)

func CheckJwt(ctx iris.Context) {
	jwt:= ctx.Request().Header["Authorization"]
	// token:= strings.Split(jwt[0], " ")[1]
	
	jwt.Verify()

	ctx.Next()
}