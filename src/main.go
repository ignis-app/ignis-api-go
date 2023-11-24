package main

import "ignis/src/handlers"
import "ignis/src/middleware"
import "context"
import "fmt"
import "os"
import "github.com/gin-gonic/gin"
import "github.com/joho/godotenv"
import "go.mongodb.org/mongo-driver/mongo"
import "go.mongodb.org/mongo-driver/mongo/options"

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
	router.Use(middleware.Auth(client))
	router.POST("/login", handlers.Login(client))
	router.POST("/logout", handlers.Logout(client))
	router.Run("localhost:8000")
}
