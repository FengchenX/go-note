package pg

import (
	"agfun/video/entity"
	"db/pg"
	"log"
	"util"
)

type SysDB struct {
	pg.SysDB
}

func (db *SysDB) CreateTable() {
	if migrate := db.AutoMigrate(entity.Video{}); migrate.Error != nil {
		log.Fatal(migrate.Error)
	}
}
func (db *SysDB) AddVideo(v *entity.Video) error {
	e := db.Create(v)
	return e.Error
}
func (db *SysDB) GetVideos(v *entity.Video, filter *util.PageFilter) ([]entity.Video, int, error) {
	vs := []entity.Video{}
	args:=db.Where(v)
	sql := util.PageFilterSql(args, "id", filter)
	if e := sql.Find(&vs); e.Error != nil {
		return nil, 0, e.Error
	}
	total := 0
	if e := args.Model(entity.Video{}).Count(&total); e.Error != nil {
		return nil, 0, e.Error
	}
	return vs, total, nil
}
