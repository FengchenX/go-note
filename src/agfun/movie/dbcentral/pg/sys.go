package pg

import (
	"agfun/movie/entity"
	"db/pg"
	"log"
	"util"
)

type SysDB struct {
	pg.SysDB
}

func (db *SysDB) CreateTable() {
	if migrate := db.AutoMigrate(entity.Movie{}, entity.MovieVideo{}); migrate.Error != nil {
		log.Fatal(migrate.Error)
	}
}

func (db *SysDB) AddMovie(m *entity.Movie) error {
	if e := db.Create(m); e.Error != nil {
		return e.Error
	}
	return nil
}

func (db *SysDB) GetMovies(m *entity.Movie, filter *util.PageFilter) ([]entity.Movie, int, error) {
	ms := []entity.Movie{}
	args := db.Where(m)
	sql := util.PageFilterSql(args, "id", filter)
	e := sql.Find(&ms)
	if e.Error != nil {
		return nil, 0, e.Error
	}
	total := 0
	e = args.Model(&entity.Movie{}).Count(&total)
	if e.Error != nil {
		return nil, 0, e.Error
	}
	return ms, total, nil
}

func (db *SysDB) AddMovieVideo(mv *entity.MovieVideo) error {
	if e := db.Create(mv); e.Error != nil {
		return e.Error
	}
	return nil
}
func (db *SysDB) GetMovieVideos(mv *entity.MovieVideo, filter *util.PageFilter) ([]entity.MovieVideo, int, error) {
	mvs := []entity.MovieVideo{}
	args := db.Where(mv)
	sql := util.PageFilterSql(args, "id", filter)
	e := sql.Find(&mvs)
	if e.Error != nil {
		return nil, 0, e.Error
	}
	total := 0
	e = args.Model(&entity.MovieVideo{}).Count(&total)
	if e.Error != nil {
		return nil, 0, e.Error
	}
	return mvs, total, nil
}
