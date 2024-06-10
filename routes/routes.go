package routes

import (
	"github.com/gauravsingh4242/bakar/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func healthCheck(ctx *gin.Context) {
	ctx.Writer.WriteHeader(200)
}

func InitRoutes(cont *controller.BakarController) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(gin.Logger())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"*"}

	router.Use(cors.New(corsConfig))

	router.GET("/bakar/health", func(c *gin.Context) { healthCheck(c) })

	group := router.Group("bakar/api/v1")

	group.GET("/ws", func(c *gin.Context) { cont.WebSocketController(c) })

	return router
}
