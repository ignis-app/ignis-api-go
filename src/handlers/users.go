package handlers

import (
	"bonfire/src/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func Users(client *mongo.Client) func (c *gin.Context) {
	return func (c *gin.Context) {
		if !c.GetBool("loggedIn") {
			c.Status(http.StatusUnauthorized)
			return
		}
		id := c.Query("id")

		var user structs.User
		coll := client.Database("bonfire").Collection("users")

		err := coll.FindOne(c, bson.M{"_id": id}).Decode(&user)
		if err == mongo.ErrNoDocuments {
			c.Status(http.StatusNotFound)
			return
		} else if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"id": user.Id,
			"username": user.Username,
			"profile": user.Profile,
			"creationdate": user.CreationDate,
		})
	}
}
