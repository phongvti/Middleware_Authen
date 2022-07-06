package user

import (
	"context"
	"time"

	"base_auth/auth"
	"base_auth/configs"
	"base_auth/models"
	"base_auth/res"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func getAll(ctx iris.Context){
	ctxx, cancel:= context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll:= configs.GetCollection("users")
	filter:= bson.M{}

	cursor, err:= coll.Find(context.TODO(), filter)
	if err!=nil{
		ctx.JSON(res.Response{Status: 500, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	defer cursor.Close(ctxx)
	var users []models.User
	for cursor.Next(ctxx){
		var user models.User
		cursor.Decode(&user)
		users= append(users, user)
	}
	ctx.JSON(res.Response{Status: 200, Message: "Success", Data: map[string]interface{}{"data": users}})
}

func register(ctx iris.Context){
	coll:= configs.GetCollection("users")
	var user models.User
	err:= ctx.ReadJSON(&user)

	if err!=nil{
		ctx.JSON(res.Response{Status: 500, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	user.Password, err= models.Hash(user.Password)

	if err!=nil{
		ctx.JSON(res.Response{Status: 404, Message: "Password can't hash", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	result, err:= coll.InsertOne(context.TODO(), user)
	if err!=nil{
		ctx.JSON(res.Response{Status: 500, Message: "ERROR", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if err!=nil{
		ctx.JSON(res.Response{Status: 500, Message: "Insert fail", Data: map[string]interface{}{"data": err.Error()}})
		return 
	}

	ctx.JSON(res.Response{Status: 201, Message: "Insert success", Data: map[string]interface{}{"data": result}})

}

func login(ctx iris.Context){
	coll:= configs.GetCollection("users")
	var body models.User
	err:= ctx.ReadJSON(&body)
	if err!=nil{
		ctx.JSON(res.Response{Status: 403, Message: "Bad request", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	var user models.User
	err= coll.FindOne(context.TODO(), bson.M{"username": body.Username}).Decode(&user)

	if err!=nil{
		ctx.JSON(res.Response{Status: 404, Message: "Incorrect username", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	err= models.CheckPasswordHash(user.Password, body.Password)
	if err!=nil{
		ctx.JSON(res.Response{Status: 401, Message: "Incorrect password", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	token, err:= auth.CreateJwt(user.Username)

	if err!=nil{
		ctx.JSON(res.Response{Status: 500, Message: "Internal server error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(bson.M{
		"id": user.ID,
		"username": user.Username,
		"role": user.Role,
		"jwt": token,
	})
}

func deleteUser (ctx iris.Context){
	coll:= configs.GetCollection("users")
	idUser:= ctx.Params().Get("idUser")
	objId, _:= primitive.ObjectIDFromHex(idUser)

	result, err:= coll.DeleteOne(context.TODO(), bson.D{{"_id", objId}})
	if err!=nil{
		ctx.JSON(res.Response{Status: 500, Message: "ERROR", Data: map[string]interface{}{"data": err.Error()}})
		return 
	}

	ctx.JSON(res.Response{Status: 200, Message: "Delete ok", Data: map[string]interface{}{"data": result}})
}