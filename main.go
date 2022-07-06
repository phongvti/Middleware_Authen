package main

import (
	"base_auth/controllers/user"

	"github.com/kataras/iris/v12"
)

func main(){
	app:= iris.New()

	user.EquipRouter(app)
	
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}