package router

import "net/http"

func initStaticFile() {
	group := rut.Group("/assets")
	group.Static("/", "./file-svc/assets")
	group.StaticFS("/more_static", http.Dir("my_file_system"))
	group.StaticFile("/favicon.ico", "./resources/favicon.ico")
}