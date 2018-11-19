package mysqldb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func Init() {
	// 初始化系统DB
	InitSysDB()

	// 初始化authDB
	InitAuthDB()
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
