package service

import (
	"agfun/agfun-service/dto"
	"agfun/tv-mgr/dbcentral/mysqldb"
	"agfun/tv-mgr/entity"
)

func (s *MovieSvc) AddFreeMovies(frees []*entity.FreeTV) error {
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
func (s *MovieSvc) GetFreeMovies(req dto.GetVideos) (*entity.GetTVResp, error) {
	movies, i, e := mysqldb.GetFreeMovies(entity.FreeTV{}, req.Filter)
	if e != nil {
		return nil, e
	}
	var resp entity.GetTVResp
	resp.Total = i
	resp.Frees = movies
	return &resp, nil
}

func (s *MovieSvc) UpdateFreeMovie(free entity.FreeTV) error {
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

func (s *MovieSvc) DelFreeMovie(free entity.FreeTV, token string) error {
	e := mysqldb.DelFreeMovie(free)
	return e
}
