package main

import (
	"base_auth/controllers"

	"github.com/kataras/iris/v12"
)

func main(){
	app:= iris.New()

	// user.EquipRouter(app)
	controllers.WithRouter(app)
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}