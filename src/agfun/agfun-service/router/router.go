package router

import (
	"agfun/agfun-service/controller"
	"agfun/agfun-service/util"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.Use(util.Cors())
	router.GET("/", controller.Hello)
	initFreeVideoRouter(router)
	return router
}
