package main

import (
	"agfun/agfun-service/dbcentral/mysqldb"
	auth "agfun/auth/dbcentral/mysqldb"
	"agfun/auth/router"
	"log"
)

func main() {
	authdb := mysqldb.InitAuthDB()
	defer authdb.Close()

	auth.CreateTable()

	r := router.Init()

	log.Fatal(r.Run(":8080"))
}
