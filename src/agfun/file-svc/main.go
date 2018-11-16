package main

import (
	"agfun/file-svc/router"
)

func main() {
	router := router.Init()

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
