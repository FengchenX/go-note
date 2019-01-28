package router

import (
	"agfun/router"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	rut = router.Init()
	initStaticFile()
	initUpload()
	return rut
}

var rut *gin.Engine
