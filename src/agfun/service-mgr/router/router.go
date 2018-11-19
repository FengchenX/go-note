package router

import "github.com/gin-gonic/gin"

var rut *gin.Engine

func Init() *gin.Engine {
	initFreeSvc()
	initPaidSvc()
	return rut
}
