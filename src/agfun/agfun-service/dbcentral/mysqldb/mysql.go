package mysqldb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func Init() {
	db, err := gorm.Open("mysql",
		"root:feng@tcp(localhost:3306)/agfun?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	sysdb = db
}

func InitSysDB() *gorm.DB {
	db, err := gorm.Open("mysql",
		"root:feng@tcp(localhost:3306)/agfun?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	sysdb = db
	return sysdb
}

var sysdb *gorm.DB

func GetSysDB() *gorm.DB {
	return sysdb
}

var authdb *gorm.DB

func GetAuthDB() *gorm.DB {
	return authdb
}

func InitAuthDB() *gorm.DB {
	db, err := gorm.Open("mysql",
		"root:feng@tcp(localhost:3306)/auth?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	authdb = db
	authdb.LogMode(true)
	return authdb
}
