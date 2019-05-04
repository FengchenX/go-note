package service

import (
	"agfun/movie/dto"
	"github.com/kataras/iris"
	"util"
)

func (s *MovieSvc) AddFree(c iris.Context) {
	req := dto.FreeMovie{}
	e := c.ReadJSON(&req)
	if e != nil {
		util.Fail(c, e)
		return
	}

}
func (s *MovieSvc) GetFrees(c iris.Context) {

}
func (s *MovieSvc) UpdateFree(c iris.Context) {

}
func (s *MovieSvc) DeleteFree(c iris.Context) {

}
