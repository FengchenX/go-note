package main

import (
	"agfun/agfun-service/dbcentral/mysqldb"
	"agfun/auth/router"
	"log"
)

func main() {
	authdb := mysqldb.InitAuthDB()
	defer authdb.Close()

	r:=router.Init()

	log.Fatal(r.Run(":8080"))
}
