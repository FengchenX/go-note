package mysqldb

import (
	sys "agfun/agfun-service/dbcentral/mysqldb"
	"database/sql"
)

func SysDB() *sql.DB {
	return sys.GetSysDB()
}

func TestDB() {
	SysDB().QueryRow("")
}
