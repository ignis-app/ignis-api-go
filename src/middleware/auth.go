package middleware

import "ignis/src/structs"
import "os"
import "github.com/gin-gonic/gin"
import "go.mongodb.org/mongo-driver/bson"
import "go.mongodb.org/mongo-driver/mongo"

// Caching might be good for this one.
func Auth(client *mongo.Client) func (c *gin.Context) {
	return func (c *gin.Context) {
		seskey, err := c.Cookie("session")
		if err != nil {
			c.Set("loggedIn", false)
			c.Next()
			return
		}
		var ses structs.Session
		coll := client.Database("ignis").Collection("sessions")
		err = coll.FindOne(c, bson.M{"key": seskey}).Decode(&ses)
		if err == mongo.ErrNoDocuments {
			c.SetCookie("session", "", -1, "/", os.Getenv("DOMAIN"), true, true)
			c.Set("loggedIn", false)
			c.Next()
			return
		} else if err != nil {
			panic(err)
		}
		var res structs.User
		coll = client.Database("ignis").Collection("users")
		if coll.FindOne(c, bson.M{"_id": ses.UserId}).Decode(&res) != nil {
			panic(err)
		}
		c.Set("loggedIn", true)
		c.Set("sessionKey", ses.Key)
		c.Set("user", res)
		c.Next()
	}	
}
