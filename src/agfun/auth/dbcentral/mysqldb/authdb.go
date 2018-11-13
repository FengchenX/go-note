package mysqldb

import (
	auth "agfun/agfun-service/dbcentral/mysqldb"
	"database/sql"
)

func AuthDB() *sql.DB {
	return auth.GetAuthDB()
}

func TestDB() {
	AuthDB().QueryRow("")
}
