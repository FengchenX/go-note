package router

import (
	"agfun/agfun-service/controller"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	//router.Use(util.Cors()) //当直接使用swagger edit 时需要打开解决跨域问题
	router.GET("/", controller.Hello)
	initFreeVideoRouter(router)
	return router
}
