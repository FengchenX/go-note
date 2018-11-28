package controller

import (
	"agfun/agfun-service/util"
	"agfun/auth/entity"
	"agfun/auth/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateUser(c *gin.Context) {
	req, err := decodeCreateUserReq(c)
	if err != nil || req == nil {
		util.Fail(c, err)
		return
	}
	user, err := service.GetDefaultSvc().CreateUser(*req)
	if err != nil {
		util.Fail(c, err)
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
		util.Fail(c, err)
		return
	}
	user, err := service.GetDefaultSvc().Login(*req)
	if err != nil {
		util.Fail(c, err)
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

func AddVip(c *gin.Context) {
	vipUser, session, e := decodeAddVipReq(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	e = service.GetDefaultSvc().AddVip(vipUser, session)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, vipUser)
}
func decodeAddVipReq(c *gin.Context) (*entity.VipUser, string, error) {
	session := c.GetHeader("auth-session")
	var vip entity.VipUser
	e := c.BindJSON(&vip)
	if e != nil {
		return nil, session, e
	}
	return &vip, session, nil
}

func GetVips(c *gin.Context) {
	vipUser, e := decodeGetVipsReq(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	users, e := service.GetDefaultSvc().GetVips(vipUser)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, users)
}
func decodeGetVipsReq(c *gin.Context) (entity.VipUser, error) {
	var vip entity.VipUser
	e := c.ShouldBindQuery(&vip)
	if e != nil {
		return vip, e
	}
	return vip, nil
}
func UpdateVip(c *gin.Context) {
	user, e := decodeUpdateVipReq(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	e = service.GetDefaultSvc().UpdateVip(user)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, nil)
}
func decodeUpdateVipReq(c *gin.Context) (*entity.VipUser, error) {
	var vip entity.VipUser
	e := c.BindJSON(&vip)
	if e != nil {
		return nil, e
	}
	return &vip, nil
}
func DelVip(c *gin.Context) {
	id, e := decodeDelVipReq(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	e = service.GetDefaultSvc().DelVip(id)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, nil)
}
func decodeDelVipReq(c *gin.Context) (int, error) {
	return strconv.Atoi(c.Param("id"))
}
