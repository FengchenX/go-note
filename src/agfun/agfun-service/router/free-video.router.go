package router

import (
	"agfun/agfun-service/controller"
	"github.com/gin-gonic/gin"
)

func initFreeVideoRouter(router *gin.Engine) {
	group := router.Group("/free-video")
	group.GET("/", controller.GetFreeVideos)
	group.POST("/", controller.AddFreeVideos)
}
