package user

import (
	"github.com/kataras/iris/v12"
	"base_auth/middlewares"
)

func EquipRouter(app iris.Party){
	userParty:= app.Party("/user")
	{
		userParty.Post("/", register)
		userParty.Post("/login", login)
		userParty.Delete("/{idUser}", deleteUser)
		

		userParty.Use(middlewares.AuthMiddleware())
		userParty.Get("/", getAll)
		
	}
}

