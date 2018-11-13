package controller

import (
	"agfun/agfun-service/util"
	"agfun/auth/entity"
	"agfun/auth/service"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	req, err := decodeCreateUserReq(c)
	if err != nil {
		util.Fail(c, nil, err)
		return
	}
	user, err := service.CreateUser(req)
	if err != nil {
		util.Fail(c, user, err)
		return
	}
	util.Success(c, user)
}

func decodeCreateUserReq(c *gin.Context) (entity.User, error) {
	var req entity.User
	if err := c.BindJSON(&req); err != nil {
		return entity.User{}, err
	}
	return req, nil
}
