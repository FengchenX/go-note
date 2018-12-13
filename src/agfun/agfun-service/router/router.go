package router

import (
	"agfun/agfun-service/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	//router.Use(util.Cors()) //当直接使用swagger edit 时需要打开解决跨域问题
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "session", "accept"}
	router.Use(cors.New(config))

	router.GET("/", controller.Hello)
	initFreeVideoRouter(router)
	return router
}
