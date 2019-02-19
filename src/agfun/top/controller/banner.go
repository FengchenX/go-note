package controller

import (
	"agfun/top/service"
	"agfun/util"
	"github.com/gin-gonic/gin"
)

func GetBanners(c *gin.Context) {
	banners, e := service.GetDefaultSvc().GetBanners()
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, banners)
}