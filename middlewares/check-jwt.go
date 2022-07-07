package middlewares

import (
	"base_auth/configs"
	"base_auth/models"
	"base_auth/res"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
)
func AuthMiddleware() context.Handler{
	return func (ctx *context.Context){
		verifier:= jwt.NewVerifier(jwt.HS256, []byte(configs.EnvSignKey()))

		verifier.WithDefaultBlocklist()
		verifyMiddleware:= verifier.Verify(func() interface{}{
			return new(models.User)
		})
		
		ok := ctx.Proceed(verifyMiddleware)
		if !ok {
			ctx.JSON(res.Response{Status: 401, Message: "Error: Unauthorized", Data: map[string]interface{}{"data": nil}})
			return
		}

		ctx.Next()
	}
}