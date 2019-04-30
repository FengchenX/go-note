package main

import (
	"agfun/movie/service"
)
func main() {


	defaultSvc := service.NewMovieSvc()
	defer defaultSvc.SysDB.Close()
	defer defaultSvc.AuthDB.Close()
	defer defaultSvc.Dynamic.Close()



	//mysqldb.CreateTable()
	//router := router.Init()
	//router.Run(":8080")
}

//package main
//
//import "github.com/kataras/iris"
//
//func main() {
//	app := iris.Default()
//	app.Get("/ping", func(ctx iris.Context) {
//		ctx.JSON(iris.Map{
//			"message": "pong",
//		})
//	})
//	// listen and serve on http://0.0.0.0:8080.
//	app.Run(iris.Addr(":8080"))
//}