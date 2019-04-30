package main

import (
	"agfun/file-svc/router"
)

func main() {
	router := router.Init()
	//router.Use(util.Cors())
	// Listen and serve on 0.0.0.0:8081
	router.Run(":8081")
}
