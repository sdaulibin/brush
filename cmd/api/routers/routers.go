package routers

import (
	"binginx.com/brush/cmd/api/handler"
	"binginx.com/brush/cmd/api/middleware/authorization"
	"binginx.com/brush/config"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() *gin.Engine {
	var router *gin.Engine
	if config.GlobalConfig.DebugMode {
		gin.SetMode(gin.DebugMode)
		router = gin.Default()
		pprof.Register(router)
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.Default()
	}
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})
	apiV1 := router.Group("/api/v1")

	apiV1.Use(authorization.CheckToken())
	{
		apiV1.GET("/userInfo", handler.UserInfo)
		apiV1.GET("/score", handler.Score)
		apiV1.GET("/news", handler.News)
		apiV1.GET("/view", handler.View)
		apiV1.GET("/behavior", handler.DoBehavior)
	}
	return router
}
