package controller

import (
	"agfun/dto"
	dto2 "agfun/movie/dto"
	"agfun/movie/service"
	"agfun/util"
	"github.com/gin-gonic/gin"
)

func GetFreeMovies(c *gin.Context) {
	req, e := decodeGetFreeMovies(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	videos, e := service.GetDefaultSvc().GetFreeMovies(req)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, videos)
}

func decodeGetFreeMovies(c *gin.Context) (dto.GetVideos, error) {
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

func AddFreeMovies(c *gin.Context) {
	videos, e := decodeCreateFreeMovies(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	for _, v := range videos {
		v.Video.ID = util.NewUUID()
		v.FreeVideo.ID = util.NewUUID()
		v.Movie.ID = util.NewUUID()
	}
	e = service.GetDefaultSvc().AddFreeMovies(videos)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, videos)
}

func decodeCreateFreeMovies(c *gin.Context) ([]*dto2.FreeMovie, error) {
	var req []*dto2.FreeMovie
	e := c.BindJSON(&req)
	return req, e
}

func UpdateFreeMovie(c *gin.Context) {
	movie, e := decodeUpdateFreeMovie(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	e = service.GetDefaultSvc().UpdateFreeMovie(*movie)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, nil)
}
func decodeUpdateFreeMovie(c *gin.Context) (*dto2.FreeMovie, error) {
	var req dto2.FreeMovie
	e := c.BindJSON(&req)
	return &req, e
}

func DelFreeMovie(c *gin.Context) {
	movie, token, e := decodeDelFreeMovie(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	e = service.GetDefaultSvc().DelFreeMovie(*movie, token)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, nil)
}
func decodeDelFreeMovie(c *gin.Context) (*dto2.FreeMovie, string, error) {
	token := c.GetHeader("session")
	var free dto2.FreeMovie
	e := c.BindJSON(&free)
	return &free, token, e
}
