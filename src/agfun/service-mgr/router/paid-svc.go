package router

import "agfun/service-mgr/controller"

func initPaidSvc() {
	group := rut.Group("/paid")
	group.GET("/", controller.GetPaidVideos)
}