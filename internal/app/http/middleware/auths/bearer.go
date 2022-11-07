package auths

import "github.com/gin-gonic/gin"

func Bearer() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.GetHeader("")
	}
}
