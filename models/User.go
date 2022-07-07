package models

import (
	"html"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)


type User struct{
	ID primitive.ObjectID 	`json:"_id" bson:"_id,omitempty"`
	Username string 		`json:"username"`
	Password string 		`json:"password"`
	Role 	 int 			`json:"role"`
}

type UserDTO struct{
	Username string 	`json:"username"`
}


func Hash (password string) (string, error){
	bytes, err:= bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}

