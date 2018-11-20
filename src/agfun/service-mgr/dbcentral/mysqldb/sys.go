package mysqldb

import (
	"agfun/agfun-service/dbcentral/mysqldb"
	"agfun/service-mgr/entity"
	"github.com/jinzhu/gorm"
)

func getSysDB() *gorm.DB{
	return mysqldb.GetSysDB()
}

func CreateTable() {
	if db := getSysDB().AutoMigrate(&entity.FreeVideo{}); db.Error != nil {
		panic(db.Error)
	}
}