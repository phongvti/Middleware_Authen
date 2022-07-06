package user

import (
	"base_auth/middlewares"

	"github.com/kataras/iris/v12"
)

func EquipRouter(app *iris.Application){
	userParty:= app.Party("/user")
	{
		userParty.Get("/", middlewares.CheckJwt, getAll)
		userParty.Post("/", register)
		userParty.Post("/login", login)
		userParty.Delete("/{idUser}", deleteUser)
	}
	
}

