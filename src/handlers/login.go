package handlers

import "ignis/src/structs"
import "ignis/src/util"
import "net/http"
import "os"
import "golang.org/x/crypto/bcrypt"
import "github.com/gin-gonic/gin"
import "go.mongodb.org/mongo-driver/bson"
import "go.mongodb.org/mongo-driver/mongo"

func Login(client *mongo.Client) func (c *gin.Context) {
	return func (c *gin.Context) {
		var body structs.Login
		if c.ShouldBindJSON(&body) != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		var user structs.User
		coll := client.Database("bonfire").Collection("users")
		
		err := coll.FindOne(c, bson.M{"email": body.Email}).Decode(&user)
		if err == mongo.ErrNoDocuments {
			c.Status(http.StatusUnauthorized)
			return
		} else if err != nil {
			panic(err)
		}
		
		var hash []byte = []byte(user.Password)
		err = bcrypt.CompareHashAndPassword(hash, []byte(body.Password))
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return
		}
		
		key := util.SessionKey()

		coll = client.Database("bonfire").Collection("sessions")
		// If workers are implemented, change the argument of Snowflake here.
		_, err = coll.InsertOne(c, bson.D{
			{Key: "_id", Value: util.Snowflake(0)},
			{Key: "userid", Value: user.Id},
			{Key: "key", Value: key},
		})
		if err != nil {
			panic(err)
		}

		c.SetCookie("session", key, 604800, "/", os.Getenv("DOMAIN"), true, true)
		c.Status(http.StatusOK)
	}
}