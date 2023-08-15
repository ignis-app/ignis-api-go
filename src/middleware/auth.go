package middleware

import (
	"bonfire/src/structs"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Caching might be good for this one.
func Auth(client *mongo.Client) func (c *gin.Context) {
	return func (c *gin.Context) {
		seskey, err := c.Cookie("session")
		if err != nil {
			c.Next()
			return
		}
		var ses structs.Session
		coll := client.Database("bonfire").Collection("sessions")
		if err := coll.FindOne(c, bson.M{"key": seskey}).Decode(&ses); err == mongo.ErrNoDocuments {
			c.SetCookie("session", "", -1, "/", os.Getenv("DOMAIN"), true, true)
			c.Next()
			return
		} else if err != nil {
			panic(err)
		}
		coll = client.Database("bonfire").Collection("users")
		
	}	
}
