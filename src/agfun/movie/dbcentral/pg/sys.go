package pg

import (
	"agfun/movie/entity"
	"db/pg"
	"log"
)

type SysDB struct {
	pg.SysDB
}

func (db *SysDB) CreateTable(){
	if migrate := db.AutoMigrate(entity.Movie{}); migrate.Error != nil {
		log.Fatal(migrate.Error)
	}
}
