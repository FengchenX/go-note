package router

import (
	"agfun/agfun-service/router"
	"github.com/gin-gonic/gin"
)

func Init() {
	rut=router.Init()
}
var rut *gin.Engine
