package mysqldb

import (
	"agfun/agfun-service/dbcentral/mysqldb"
	"github.com/jinzhu/gorm"
)

func getSysDB() *gorm.DB{
	return mysqldb.GetSysDB()
}
