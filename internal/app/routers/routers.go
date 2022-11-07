package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/pro911/authApis/internal/app/http/controller"
	"github.com/pro911/authApis/internal/app/http/middleware"
	"github.com/pro911/authApis/internal/app/http/middleware/auths"
	"net/http"
)

func RegisterRouter(r *gin.Engine) {

	//解决跨域中间件
	r.Use(middleware.Cors())

	//一个测试监听探活接口
	r.Any("online", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
		return
	})

	auth := r.Group("auth")

	auth.Any("ping_kv", auths.Kv(), controller.Ping)

	auth.Any("ping_bearer", auths.Bearer(), controller.Ping)

	auth.Any("ping_basic", auths.Basic(), controller.Ping)
}
