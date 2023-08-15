package main

import (
	"bonfire/src/handlers"

	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No environment file found.")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		panic("Mongo URI not set.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	router := gin.Default()

	router.POST("/login", handlers.Login(client))
	router.POST("/logout", handlers.Logout(client))
	router.Run("localhost:8000")
}
