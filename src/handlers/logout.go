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
		coll := client.Database("bonfire").Collection("sessions")
		cookie, err := c.Request.Cookie("session")
		if err != nil {
			panic(err)
		}
		if _, err := coll.DeleteOne(c, bson.M{"key": cookie.Value}); err != nil {
			panic(err)
		}
		c.SetCookie("session", "", -1, "/", os.Getenv("DOMAIN"), true, true)
		c.Status(http.StatusAccepted)
	}
}
