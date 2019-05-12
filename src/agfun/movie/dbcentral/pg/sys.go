package pg

import (
	"agfun/movie/entity"
	"db/pg"
	"log"
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

func (db *SysDB) GetMovies(m entity.Movie) ([]entity.Movie, int, error) {
	ms := []entity.Movie{}
	e := db.Find(&ms).Where(&m)
	if e.Error != nil {
		return nil, 0, e.Error
	}
	total := 0
	e = db.Count(&total).Where(&m)
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
