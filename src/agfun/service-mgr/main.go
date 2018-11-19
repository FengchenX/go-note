package main

import "agfun/service-mgr/router"

func main() {

	router := router.Init()
	router.Run("8080")
}
