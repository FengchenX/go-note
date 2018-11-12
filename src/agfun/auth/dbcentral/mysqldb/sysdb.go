package mysqldb

import (
	sys "agfun/agfun-service/dbcentral/mysqldb"
	"database/sql"
)

func AuthDB() *sql.DB {
	return sys.GetAuthDB()
}

func TestDB() {
	AuthDB().QueryRow("")
}
