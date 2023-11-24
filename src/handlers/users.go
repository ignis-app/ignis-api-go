package handlers

import "ignis/src/structs"
import "net/http"
import "github.com/gin-gonic/gin"
import "go.mongodb.org/mongo-driver/mongo"
import "go.mongodb.org/mongo-driver/bson"

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
