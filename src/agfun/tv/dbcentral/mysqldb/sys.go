package mysqldb

import (
	"agfun/dbcentral/mysqldb"
	"agfun/tv/entity"
	"github.com/jinzhu/gorm"
)

func getSysDB() *gorm.DB {
	return mysqldb.GetSysDB()
}

func CreateTable() {
	if db := getSysDB().AutoMigrate(&entity.FreeTV{}); db.Error != nil {
		panic(db.Error)
	}
}
