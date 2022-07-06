package configs

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func EnvMongoURI() string{
	err:= godotenv.Load()
	if err!=nil{
		log.Fatal("Error load .env file")
	}

	return os.Getenv("MONGO_URI")
}

func EnvSecretKey() string{
	err:= godotenv.Load()
	if err!=nil{
		log.Fatal("Error load .env file")
	}

	return os.Getenv("SECRET_JWT")
}