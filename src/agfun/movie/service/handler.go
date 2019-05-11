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
	util.Success(c, &m)
}
func (s *MovieSvc) GetMovies(c iris.Context) {

}
func (s *MovieSvc) UpdateMovie(c iris.Context) {

}
func (s *MovieSvc) DeleteMovie(c iris.Context) {

}
