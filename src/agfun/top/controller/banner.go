package controller

import (
	"agfun/dto"
	"agfun/top/service"
	"agfun/util"
	"github.com/gin-gonic/gin"
)

func GetBanners(c *gin.Context) {
	videos, e := service.GetDefaultSvc().GetFreeTVs(req)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, videos)
}