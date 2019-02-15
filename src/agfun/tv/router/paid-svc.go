package router

import "agfun/tv/controller"

func initPaidSvc() {
	group := rut.Group("/paid")
	group.GET("/", controller.GetPaidVideos)
}
