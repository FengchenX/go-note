package router

import (
	"agfun/agfun-service/router"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	rut = router.Init()
	initUser()
	return rut
}

var rut *gin.Engine
