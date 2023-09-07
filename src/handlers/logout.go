package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Logout(client *mongo.Client) func (c *gin.Context) {
	return func (c *gin.Context) {
		if !c.GetBool("loggedIn") {
			c.Status(http.StatusAccepted)
			return
		}
		coll := client.Database("bonfire").Collection("sessions")
		if _, err := coll.DeleteOne(c, bson.M{"key": c.GetString("sessionKey")}); err != nil {
			panic(err)
		}
		c.SetCookie("session", "", -1, "/", os.Getenv("DOMAIN"), true, true)
		c.Status(http.StatusAccepted)
	}
}
