package controller

import (
	"agfun/agfun-service/util"
	"agfun/movie-mgr/dto"
	"agfun/movie-mgr/entity"
	"agfun/movie-mgr/service"
	"github.com/gin-gonic/gin"
)

func GetFreeMovies(c *gin.Context) {
	req, e := decodeGetFreeMovies(c)
	if e != nil {
		util.Fail(c, nil, e)
		return
	}
	videos, e := service.GetDefaultSvc().GetFreeMovies(req)
	if e != nil {
		util.Fail(c, nil, e)
		return
	}
	util.Success(c, videos)
}

func decodeGetFreeMovies(c *gin.Context) (dto.GetFreeMovies, error) {
	var req dto.GetFreeMovies
	filter, e := util.ParsePageFilter(c)
	if e != nil {
		return req, e
	}
	token := c.GetHeader("auth-session")
	req.Filter = filter
	req.Token = token
	return req, nil
}

func AddFreeMovies(c *gin.Context) {
	videos, e := decodeCreateFreeMovies(c)
	if e != nil {
		util.Fail(c, nil, e)
		return
	}
	e = service.GetDefaultSvc().AddFreeMovies(videos)
	if e != nil {
		util.Fail(c, nil, e)
		return
	}
	util.Success(c, videos)
}

func decodeCreateFreeMovies(c *gin.Context) ([]*entity.FreeMovie, error) {
	var req []*entity.FreeMovie
	e := c.BindJSON(&req)
	return req, e
}
