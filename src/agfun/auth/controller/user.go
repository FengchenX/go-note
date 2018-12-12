package controller

import (
	"agfun/agfun-service/util"
	"agfun/auth/dto"
	"agfun/auth/entity"
	"agfun/auth/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func CreateUser(c *gin.Context) {
	req, err := decodeCreateUserReq(c)
	if err != nil || req == nil {
		util.Fail(c, err)
		return
	}
	req.ID = util.NewUUID()
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
		util.Fail(c, fmt.Errorf("req is error"))
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
	pwd, b := c.GetQuery("pwd")
	if b {
		user.Pwd = pwd
	}
	if user.UserName == "" || user.Pwd == "" {
		return nil, fmt.Errorf("user name or pwd is nil")
	}
	return user, nil
}

func AddUserRole(c *gin.Context) {
	vipUser, session, e := decodeAddUserRoleReq(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	vipUser.ID = util.NewUUID()
	e = service.GetDefaultSvc().AddVip(vipUser, session)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, vipUser)
}
func decodeAddUserRoleReq(c *gin.Context) (*entity.UserRole, string, error) {
	session := c.GetHeader("auth-session")
	var userRole dto.UserRole
	e := c.BindJSON(&userRole)
	if e != nil {
		return nil, session, e
	}
	userRole.UserRole.Expire, e = time.Parse("2006-01-02 15:04:05", userRole.Expire)
	if e != nil {
		return nil, session, e
	}
	return &userRole.UserRole, session, nil
}

func GetVips(c *gin.Context) {
	vipUser, e := decodeGetVipsReq(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	vipUser.Expire = time.Now()
	users, e := service.GetDefaultSvc().GetVips(vipUser)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, users)
}
func decodeGetVipsReq(c *gin.Context) (entity.UserRole, error) {
	var vip entity.UserRole
	level, b := c.GetQuery("level")
	if b {
		i, e := strconv.Atoi(level)
		if e != nil {
			return vip, e
		}
		vip.Level = i
	}
	user_id, b := c.GetQuery("user_id")
	if b {
		vip.UserID = user_id
	}
	id, b := c.GetQuery("id")
	if b {
		vip.ID = id
	}
	return vip, nil
}
func UpdateVip(c *gin.Context) {
	user, session, e := decodeUpdateVipReq(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	e = service.GetDefaultSvc().UpdateVip(user, session)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, nil)
}
func decodeUpdateVipReq(c *gin.Context) (*entity.UserRole, string, error) {
	var vip dto.UserRole
	e := c.BindJSON(&vip)
	if e != nil {
		return nil, "", e
	}
	id := c.Param("id")
	vip.ID = id
	vip.UserRole.Expire, e = time.Parse("2006-01-02 15:04:05", vip.Expire)
	if e != nil {
		return nil, "", e
	}
	session := c.GetHeader("auth-session")
	return &vip.UserRole, session, nil
}
func DelVip(c *gin.Context) {
	id, session := decodeDelVipReq(c)
	e := service.GetDefaultSvc().DelVip(id, session)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, nil)
}
func decodeDelVipReq(c *gin.Context) (id string, session string) {
	session = c.GetHeader("auth-session")
	id = c.Param("id")
	return id, session
}

func AddRole(c *gin.Context) {

}
func DelRole(c *gin.Context) {

}
func UpdateRole(c *gin.Context) {

}
func GetRoles(c *gin.Context) {

}
