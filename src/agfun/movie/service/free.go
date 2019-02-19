package service

import (
	"agfun/dto"
	"agfun/movie/dbcentral/mysqldb"
	"agfun/movie/entity"
)

func (s *MovieSvc) AddFreeMovies(frees []*entity.FreeMovie) error {
	for _, free := range frees {
		_, i, e := mysqldb.GetFreeMovies(*free, nil)
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
func (s *MovieSvc) GetFreeMovies(req dto.GetVideos) (*entity.GetMoviesResp, error) {
	movies, i, e := mysqldb.GetFreeMovies(entity.FreeMovie{}, req.Filter)
	if e != nil {
		return nil, e
	}
	var resp entity.GetMoviesResp
	resp.Total = i
	resp.Frees = movies
	return &resp, nil
}

func (s *MovieSvc) UpdateFreeMovie(free entity.FreeMovie) error {
	if len(free.ID) > 0 {
		e := mysqldb.UpdateFreeMovie(free, nil)
		if e != nil {
			return e
		}
		return nil
	}
	//todo
	return nil
}

func (s *MovieSvc) DelFreeMovie(free entity.FreeMovie, token string) error {
	e := mysqldb.DelFreeMovie(free)
	return e
}