package permit

import "net/http"

import "github.com/gin-gonic/gin"
import "go.mongodb.org/mongo-driver/mongo"

func LoggedIn(client *mongo.Client) gin.HandlerFunc {
	return func (c *gin.Context) {
		if !c.GetBool("loggedIn") {
			c.Status(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
