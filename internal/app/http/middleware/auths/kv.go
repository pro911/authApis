package auths

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Kv() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		fmt.Println(token)
		if token == "apipost" {
			c.Set("userId", "123456")
			c.Next()
		} else {
			c.Abort() //阻止调用后续的处理函数
			c.JSON(http.StatusOK, gin.H{
				"errCode": 10001,
				"errMsg":  "验证失败!",
			})
			return
		}
	}
}
