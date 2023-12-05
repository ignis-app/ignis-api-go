package main

import "context"
import "fmt"
import "ignis/src/handlers"
import "ignis/src/middleware"
import "ignis/src/middleware/permit"
import "os"

import "github.com/gin-gonic/gin"
import "github.com/joho/godotenv"
import "go.mongodb.org/mongo-driver/mongo"
import "go.mongodb.org/mongo-driver/mongo/options"

func main() {
	if godotenv.Load() != nil {
		fmt.Println("No environment file found.")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		panic("Mongo URI not set.")
	}

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if client.Disconnect(ctx) != nil {
			panic(err)
		}
	}()

	router := gin.Default()
	router.Use(middleware.Auth(client))
	router.POST("/login", handlers.Login(client))
	router.POST("/logout", handlers.Logout(client))
	router.GET("/users", permit.LoggedIn(client), handlers.User(client))
	router.Run("localhost:8000")
}
