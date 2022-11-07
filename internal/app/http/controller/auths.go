package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 10000,
		"msg":  "success",
		"data": nil,
	})
	return
}
