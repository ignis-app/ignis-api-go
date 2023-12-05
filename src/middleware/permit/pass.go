package permit

import "github.com/gin-gonic/gin"
import "go.mongodb.org/mongo-driver/mongo"

func Pass(client *mongo.Client, permfn ...func (c *gin.Context) bool) gin.HandlerFunc {
	return func (c *gin.Context) {
		for _, fn := range permfn {
			if !fn(c) {
				return
			}
		}
		c.Next()
	}
}
