package main

import (
	"agfun/agfun-service/dbcentral/mysqldb"
	"agfun/auth/dbcentral/etcd"
	auth "agfun/auth/dbcentral/mysqldb"
	"agfun/auth/router"
	"log"
)

func main() {
	authdb := mysqldb.InitAuthDB()
	defer authdb.Close()

	auth.CreateTable()

	client := etcd.Init()
	defer client.Close()

	r := router.Init()

	log.Fatal(r.Run(":8080"))
}
