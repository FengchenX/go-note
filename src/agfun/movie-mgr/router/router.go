package router

import (
	"agfun/agfun-service/router"
	"github.com/gin-gonic/gin"
)

var rut *gin.Engine

func Init() *gin.Engine {
	rut = router.Init()
	initFreeSvc()
	initPaidSvc()
	return rut
}
