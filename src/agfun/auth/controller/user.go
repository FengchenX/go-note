package controller

import (
	"agfun/agfun-service/util"
	"agfun/auth/protocol"
	"agfun/auth/service"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	req, _ := decodeCreateUserReq(c)
	user, err := service.CreateUser(req)
	if err != nil {
		util.Fail(c, user, err)
	}
	util.Success(c, user)
}

func decodeCreateUserReq(c *gin.Context) (protocol.CreateUserReq, error) {
	panic("todo")
}
