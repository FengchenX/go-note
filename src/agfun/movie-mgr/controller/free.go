package controller

import (
	"agfun/agfun-service/util"
	"agfun/movie-mgr/dto"
	"agfun/movie-mgr/entity"
	"agfun/movie-mgr/service"
	"github.com/gin-gonic/gin"
)

func GetFreeVideos(c *gin.Context) {
	req, e := decodeGetFreeVideos(c)
	if e != nil {
		util.Fail(c, nil, e)
		return
	}
	videos, e := service.GetDefaultSvc().GetFreeVideos(req)
	if e != nil {
		util.Fail(c, nil, e)
		return
	}
	util.Success(c, videos)
}

func decodeGetFreeVideos(c *gin.Context) (dto.GetFreeVideos, error) {
	var req dto.GetFreeVideos
	filter, e := util.ParsePageFilter(c)
	if e != nil {
		return req, e
	}
	token := c.GetHeader("auth-session")
	req.Filter = filter
	req.Token = token
	return req, nil
}

func AddFreeVideos(c *gin.Context) {
	videos, e := decodeCreateFreeVideos(c)
	if e != nil {
		util.Fail(c, nil, e)
		return
	}
	e = service.GetDefaultSvc().AddFreeVideos(videos)
	if e != nil {
		util.Fail(c, nil, e)
		return
	}
	util.Success(c, videos)
}

func decodeCreateFreeVideos(c *gin.Context) ([]*entity.FreeVideo, error) {
	var req []*entity.FreeVideo
	e := c.BindJSON(&req)
	return req, e
}
