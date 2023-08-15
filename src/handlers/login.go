package handlers

import (
	"bonfire/src/bindings"
	"bonfire/src/util"

	"net/http"
	"os"
	"crypto/rand"
	"encoding/base64"
	"math/big"

	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(client *mongo.Client) func (c *gin.Context) {
	return func (c *gin.Context) {
		var body bindings.LoginR
		if err := c.ShouldBindJSON(&body); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		var result bson.M
		coll := client.Database("bonfire").Collection("users")
		
		if err := coll.FindOne(c, bson.M{"email": body.Email}).Decode(&result); err == mongo.ErrNoDocuments {
			c.Status(http.StatusUnauthorized)
			return
		} else if err != nil {
			panic(err)
		}
		
		var hash []byte = []byte(result["password"].(string))
		if err := bcrypt.CompareHashAndPassword(hash, []byte(body.Password)); err != nil {
			c.Status(http.StatusUnauthorized)
			return
		}

		keyInt, keyErr := rand.Int(rand.Reader, big.NewInt(4294967296))
		if keyErr != nil {
			panic(keyErr)
		}
		key := base64.RawURLEncoding.EncodeToString(keyInt.Bytes())

		// If workers are implemented, change the argument of Snowflake here.
		_, err := coll.InsertOne(c, bson.D{
			{Key: "_id", Value: util.Snowflake(0)},
			{Key: "userid", Value: result["id"].(int)},
			{Key: "key", Value: key},
		})
		if err != nil {
			panic(err)
		}

		c.SetCookie("session", key, 604800, "/", os.Getenv("DOMAIN"), true, true)
		c.Status(http.StatusAccepted)
	}
}