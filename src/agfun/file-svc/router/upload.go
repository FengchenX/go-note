package router

import "agfun/file-svc/controller"

func initUpload() {
	group := rut.Group("/upload")
	group.POST("/", controller.Test)
}
