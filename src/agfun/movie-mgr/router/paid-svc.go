package router

import "agfun/movie-mgr/controller"

func initPaidSvc() {
	group := rut.Group("/paid")
	group.GET("/", controller.GetPaidVideos)
}
