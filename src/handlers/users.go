package handlers

import (
	"bonfire/src/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func get()

func Users(client *mongo.Client) func (c *gin.Context) {
	return func (c *gin.Context) {
		if !c.GetBool("loggedIn") {
			c.Status(http.StatusUnauthorized)
			return
		}
		coll := client.Database("bonfire").Collection("users")
		var user structs.User
		if _, err := coll.FindOne(c, bson.M{"_id", })
	}
}