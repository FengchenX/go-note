package router

import "agfun/service-mgr/controller"

func initFreeSvc() {
	group := rut.Group("/free")
	group.GET("/", controller.GetFreeVideos)
	group.POST("/", controller.CreateFreeVideos)
}
