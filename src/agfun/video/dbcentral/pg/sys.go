package pg

import (
	"db/pg"
	"entity"
	"log"
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
func (db *SysDB) GetVideos(v *entity.Video) ([]entity.Video, int, error) {
	vs := []entity.Video{}
	if e := db.Find(&vs).Where(v); e.Error != nil {
		return nil, 0, e.Error
	}
	total := 0
	if e := db.Model(&entity.Video{}).Where(v).Count(&total); e.Error != nil {
		return nil, 0, e.Error
	}
	return vs, total, nil
}
