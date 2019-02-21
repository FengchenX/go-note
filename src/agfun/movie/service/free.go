package service

import (
	mysqldb2 "agfun/dbcentral/mysqldb"
	"agfun/dto"
	entity2 "agfun/entity"
	"agfun/movie/dbcentral/mysqldb"
	dto2 "agfun/movie/dto"
	"agfun/movie/entity"
	"agfun/util"
	"fmt"
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
func (s *MovieSvc) GetFreeMovies(req dto.GetVideos) (*dto2.GetMoviesResp, error) {
	movies, i, e := mysqldb.GetFreeMovies(entity.FreeMovie{}, req.Filter)
	if e != nil {
		return nil, e
	}
	var resp dto2.GetMoviesResp
	resp.Total = i
	for _, movie := range movies {
		arg := entity.Movie{
			ID: movie.MovieID,
		}
		movies, total, e := mysqldb.GetMovies(arg, nil)
		if e != nil {
			return &dto2.GetMoviesResp{}, e
		}
		if total != 1 {
			return &dto2.GetMoviesResp{}, fmt.Errorf("total != 1")
		}
		var freeMovie dto2.FreeMovie
		e = util.Copy(&freeMovie.Movie, &movies[0])
		if e != nil {
			return &dto2.GetMoviesResp{}, e
		}
		arg2 := entity2.FreeVideo{
			ID: movie.FreeVideoID,
		}
		freeVideos, total, e := mysqldb2.GetFreeVideos(arg2, nil)
		if e != nil {
			return &dto2.GetMoviesResp{}, e
		}
		if total != 1 {
			return &dto2.GetMoviesResp{}, fmt.Errorf("total != 1")
		}
		arg3 := entity2.Video{
			ID: freeVideos[0].ID,
		}
		util.Copy(&freeMovie.FreeVideo, freeVideos[0])
		videos, total, e := mysqldb2.GetVideos(arg3, nil)
		if e != nil {
			return &dto2.GetMoviesResp{}, e
		}
		if total != 1 {
			return &dto2.GetMoviesResp{}, fmt.Errorf("total != 1")
		}
		util.Copy(&freeMovie.FreeVideo, videos[0])
		resp.Frees = append(resp.Frees, &freeMovie)
	}
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
