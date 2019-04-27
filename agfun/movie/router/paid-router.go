package router

import "agfun/movie/controller"

func initPaidSvc() {
	group := rut.Group("/paid")
	group.GET("/", controller.GetPaidVideos)
}
