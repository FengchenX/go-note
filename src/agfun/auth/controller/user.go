package controller

import (
	"agfun/agfun-service/util"
	"agfun/auth/entity"
	"agfun/auth/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	req, err := decodeCreateUserReq(c)
	if err != nil || req == nil {
		util.Fail(c, nil, err)
		return
	}
	user, err := service.CreateUser(*req)
	if err != nil {
		util.Fail(c, user, err)
		return
	}
	util.Success(c, user)
}

func decodeCreateUserReq(c *gin.Context) (*entity.User, error) {
	var req entity.User
	if err := c.BindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func Login(c *gin.Context) {
	req, err := decodeLoginReq(c)
	if err != nil || req == nil {
		util.Fail(c, nil, err)
		return
	}
	user, err := service.Login(*req)
	if err != nil {
		util.Fail(c, user, err)
		return
	}
	util.Success(c, user)
}

func decodeLoginReq(c *gin.Context) (*entity.User, error) {
	user := &entity.User{
		UserName: c.Param("user-name"),
		Pwd:      c.Param("pwd"),
	}
	if user.UserName == "" || user.Pwd == "" {
		return nil, fmt.Errorf("user name or pwd is nil")
	}
	return user, nil
}
