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
		UserName: c.Query("user-name"),
		Pwd:      c.Query("pwd"),
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
	e = service.GetDefaultSvc().AddRole(vipUser, session)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, vipUser)
}
func decodeAddUserRoleReq(c *gin.Context) (*entity.UserRole, string, error) {
	session := c.GetHeader("session")

	var userRole dto.UserRole
	e := c.BindJSON(&userRole)
	if e != nil {
		return nil, session, e
	}
	userId := c.Param("user-id")
	roleId := c.Param("role-id")
	userRole.UserID = userId
	userRole.RoleID = roleId
	userRole.UserRole.Expire, e = time.Parse("2006-01-02 15:04:05", userRole.Expire)
	if e != nil {
		return nil, session, e
	}
	return &userRole.UserRole, session, nil
}

func GetUserRoles(c *gin.Context) {
	vipUser, e := decodeGetUserRolesReq(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	vipUser.Expire = time.Now()
	users, e := service.GetDefaultSvc().GetUserRoles(vipUser)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, users)
}
func decodeGetUserRolesReq(c *gin.Context) (entity.UserRole, error) {
	var vip entity.UserRole
	level, b := c.GetQuery("level")
	if b {
		i, e := strconv.Atoi(level)
		if e != nil {
			return vip, e
		}
		vip.Level = i
	}
	userIid := c.Param("user-id")
	vip.UserID = userIid
	return vip, nil
}
func UpdateUserRole(c *gin.Context) {
	user, session, e := decodeUpdateUserRoleReq(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	e = service.GetDefaultSvc().UpdateUserRole(user, session)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, nil)
}
func decodeUpdateUserRoleReq(c *gin.Context) (*entity.UserRole, string, error) {
	var vip dto.UserRole
	e := c.BindJSON(&vip)
	if e != nil {
		return nil, "", e
	}
	userId := c.Param("user-id")
	roleId := c.Param("role-id")
	vip.UserID, vip.RoleID = userId, roleId
	if len(vip.Expire) > 0 {
		vip.UserRole.Expire, e = time.Parse("2006-01-02 15:04:05", vip.Expire)
		if e != nil {
			return nil, "", e
		}
	}
	session := c.GetHeader("session")
	return &vip.UserRole, session, nil
}
func DelUserRole(c *gin.Context) {
	userRole, session, e := decodeDelUserRoleReq(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	e = service.GetDefaultSvc().DelUserRole(userRole, session)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, nil)
}
func decodeDelUserRoleReq(c *gin.Context) (userRole entity.UserRole, session string, e error) {
	session = c.GetHeader("session")
	var vip entity.UserRole
	level, b := c.GetQuery("level")
	if b {
		i, e := strconv.Atoi(level)
		if e != nil {
			return vip, "", e
		}
		vip.Level = i
	}
	vip.UserID = c.Param("user-id")
	vip.RoleID = c.Param("role-id")
	return userRole, session, nil
}

func AddRole(c *gin.Context) {

}
func DelRole(c *gin.Context) {

}
func UpdateRole(c *gin.Context) {

}
func GetRoles(c *gin.Context) {

}

func GetResources(c *gin.Context) {
	resources, e := service.GetDefaultSvc().GetResources()
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, resources)
}
