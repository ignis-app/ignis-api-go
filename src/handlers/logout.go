package handlers

import "net/http"
import "os"
import "github.com/gin-gonic/gin"
import "go.mongodb.org/mongo-driver/bson"
import "go.mongodb.org/mongo-driver/mongo"

func Logout(client *mongo.Client) gin.HandlerFunc {
	return func (c *gin.Context) {
		if !c.GetBool("loggedIn") {
			c.Status(http.StatusOK)
			return
		}
		coll := client.Database("ignis").Collection("sessions")
		_, err := coll.DeleteOne(c, bson.M{"key": c.GetString("sessionKey")})
		if err != nil {
			panic(err)
		}
		c.SetCookie("session", "", -1, "/", os.Getenv("DOMAIN"), true, true)
		c.Status(http.StatusOK)
	}
}
