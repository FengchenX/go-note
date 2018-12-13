package controller

import (
	"agfun/agfun-service/dto"
	"agfun/agfun-service/util"
	"agfun/tv-mgr/entity"
	"agfun/tv-mgr/service"
	"github.com/gin-gonic/gin"
)

func GetFreeTVs(c *gin.Context) {
	req, e := decodeGetFreeTV(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	videos, e := service.GetDefaultSvc().GetFreeTVs(req)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, videos)
}

func decodeGetFreeTV(c *gin.Context) (dto.GetVideos, error) {
	var req dto.GetVideos
	filter, e := util.ParsePageFilter(c)
	if e != nil {
		return req, e
	}
	token := c.GetHeader("session")
	req.Filter = filter
	req.Token = token
	return req, nil
}

func AddFreeTVs(c *gin.Context) {
	videos, e := decodeCreateFreeTV(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	for _, v := range videos {
		v.ID = util.NewUUID()
	}
	e = service.GetDefaultSvc().AddFreeTVs(videos)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, videos)
}

func decodeCreateFreeTV(c *gin.Context) ([]*entity.FreeTV, error) {
	var req []*entity.FreeTV
	e := c.BindJSON(&req)
	return req, e
}

func UpdateFreeTV(c *gin.Context) {
	movie, e := decodeUpdateFreeTV(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	e = service.GetDefaultSvc().UpdateFreeTV(*movie)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, nil)
}
func decodeUpdateFreeTV(c *gin.Context) (*entity.FreeTV, error) {
	var req entity.FreeTV
	e := c.BindJSON(&req)
	return &req, e
}

func DelFreeTV(c *gin.Context) {
	movie, token, e := decodeDelFreeTV(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	e = service.GetDefaultSvc().DelFreeTV(*movie, token)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, nil)
}
func decodeDelFreeTV(c *gin.Context) (*entity.FreeTV, string, error) {
	token := c.GetHeader("session")
	var free entity.FreeTV
	e := c.BindJSON(&free)
	return &free, token, e
}
