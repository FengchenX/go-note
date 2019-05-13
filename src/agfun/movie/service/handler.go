package service

import (
	"agfun/movie/dto"
	"agfun/movie/entity"
	"encoding/json"
	"github.com/kataras/iris"
	"util"
)

func (s *MovieSvc) AddMovie(c iris.Context) {
	req := dto.Movie{}
	e := c.ReadJSON(&req)
	if e != nil {
		util.Fail(c, e)
		return
	}
	m := entity.Movie{}
	util.Copy(&m, &req.Movie)
	m.ID = util.NewUUID()
	m.CreateAt = util.TimeNowStd()
	if len(req.MainPlayers) > 0 {
		buf, e := json.Marshal(req.MainPlayers)
		if e != nil {
			util.Fail(c, e)
			return
		}
		m.MainPlayers = string(buf)
	}

	if len(req.Types) > 0 {
		buf, e := json.Marshal(req.Types)
		if e != nil {
			util.Fail(c, e)
			return
		}
		m.Types = string(buf)
	}
	e = s.SysDB.AddMovie(&m)
	if e != nil {
		util.Fail(c, e)
		return
	}
	res := dto.Movie{}
	util.Copy(&res.Movie, &m)
	res.MainPlayers = req.MainPlayers
	res.Types = req.Types
	util.Success(c, &res)
}
func (s *MovieSvc) GetMovies(c iris.Context) {
	filter := util.ParsePageFilter(c)
	movies, i, e := s.SysDB.GetMovies(nil, filter)
	if e != nil {
		util.Fail(c, e)
		return
	}
	res := dto.Movies{
		Total:  i,
		Movies: nil,
	}
	for _, movie:=range movies{
		dtoMovie:=dto.Movie{
			Movie:       movie,
			MainPlayers: nil,
			Types:       nil,
		}
		json.Unmarshal([]byte(movie.MainPlayers), &dtoMovie.MainPlayers)
		json.Unmarshal([]byte(movie.Types), &dtoMovie.Types)
		res.Movies = append(res.Movies, dtoMovie)
	}
	util.Success(c, &res)
}
func (s *MovieSvc) GetMovie(c iris.Context) {

}
func (s *MovieSvc) UpdateMovie(c iris.Context) {

}
func (s *MovieSvc) DeleteMovie(c iris.Context) {

}
func (s *MovieSvc) AddMovieVideo(c iris.Context) {
	req := dto.MovieVideo{}
	e := c.ReadJSON(&req)
	if e != nil {
		util.Fail(c, e)
		return
	}
	mv := entity.MovieVideo{}
	util.Copy(&mv, &req.MovieVideo)
	mv.ID = util.NewUUID()
	mv.CreateAt = util.TimeNowStd()
	mv.MovieID = c.Params().Get("id")
	e = s.SysDB.AddMovieVideo(&mv)
	if e != nil {
		util.Fail(c, e)
		return
	}
	res := dto.MovieVideo{}
	util.Copy(&res.MovieVideo, &mv)
	util.Success(c, &res)
}
