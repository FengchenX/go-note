package router

import "agfun/tv-mgr/controller"

func initPaidSvc() {
	group := rut.Group("/paid")
	group.GET("/", controller.GetPaidVideos)
}
