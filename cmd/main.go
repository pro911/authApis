package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pro911/authApis/internal/app/routers"
	"github.com/pro911/authApis/internal/pkg"
)

func init() {
	pkg.InitProjects()
}

func main() {
	//创建路由
	r := gin.New()

	routers.RegisterRouter(r)

	if err := r.Run(":8001"); err != nil {
		panic(fmt.Errorf("httpServer:8001 run fail:%w", err))
	}
}
