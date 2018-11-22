package router

import "agfun/movie-mgr/controller"

func initFreeSvc() {
	group := rut.Group("/free")
	group.GET("/", controller.GetFreeVideos)
	group.POST("/", controller.AddFreeVideos)
}
