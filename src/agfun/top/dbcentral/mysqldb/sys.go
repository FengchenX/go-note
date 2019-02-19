package mysqldb

import (
	"agfun/dbcentral/mysqldb"
	"agfun/top/entity"
	"github.com/jinzhu/gorm"
)

func getSysDB() *gorm.DB {
	return mysqldb.GetSysDB()
}

func CreateTable() {
	if db := getSysDB().AutoMigrate(&entity.Banner{}); db.Error != nil {
		panic(db.Error)
	}
}
