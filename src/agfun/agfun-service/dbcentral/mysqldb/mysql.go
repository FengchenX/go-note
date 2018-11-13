package mysqldb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Init() {
	db, err := sql.Open("mysql",
		"feng:feng@tcp(localhost:3306)/agfun?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	sysdb = db
}

func InitSysDB() *sql.DB {
	db, err := sql.Open("mysql",
		"feng:feng@tcp(localhost:3306)/agfun?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	sysdb = db
	return sysdb
}

var sysdb *sql.DB
func GetSysDB() *sql.DB {
	return sysdb
}

var authdb *sql.DB
func GetAuthDB() *sql.DB{
	return authdb
}

func InitAuthDB() *sql.DB{
	db, err := sql.Open("mysql",
		"feng:feng@tcp(localhost:3306)/auth?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	sysdb = db
	return sysdb
}
