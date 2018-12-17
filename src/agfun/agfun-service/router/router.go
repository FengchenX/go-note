package router

import (
	"agfun/agfun-service/controller"
	"agfun/agfun-service/dbcentral/etcddb"
	"agfun/agfun-service/dbcentral/mysqldb"
	"agfun/agfun-service/jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, []string{"session", "accept"}...)
	router.Use(cors.New(config), jwt.AuthMiddleWare(mysqldb.GetAuthDB(), etcddb.GetCli()))

	router.GET("/", controller.Hello)
	initFreeVideoRouter(router)
	return router
}
