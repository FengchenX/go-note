package router

import (
	"agfun/agfun-service/controller"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.GET("/", controller.Hello)
	initFreeVideoRouter(router)
	return router
}
