package router

import (
	"agfun/router"
	"github.com/gin-gonic/gin"
)

var rut *gin.Engine
var rutGroup *gin.RouterGroup

func Init() *gin.Engine {
	rut = router.Init()
	rutGroup = rut.Group("/top")

	initBannerSvc()
	return rut
}
