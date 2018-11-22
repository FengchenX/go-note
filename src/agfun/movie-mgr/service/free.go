package service

import (
	"agfun/movie-mgr/dbcentral/mysqldb"
	"agfun/movie-mgr/dto"
	"agfun/movie-mgr/entity"
)

func (s *MgrSvc) AddFreeMovies(frees []*entity.FreeMovie) error {
	for _, free := range frees {
		_, i, e := mysqldb.GetFreeVideos(*free, nil)
		if e != nil {
			return e
		}
		if i > 0 {
			continue
		}
		e = mysqldb.AddFreeMovies(free)
		if e != nil {
			return e
		}
	}
	return nil
}
func (s *MgrSvc) GetFreeMovies(req dto.GetFreeMovies) ([]*entity.FreeMovie, error) {
	var id int
	e := s.Dynamic.Get(req.Token, &id)
	if e != nil {
		return nil, e
	}
	var videos []*entity.FreeMovie
	return videos, nil
}
