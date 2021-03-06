package service

import (
	"agfun/auth/dbcentral/mysqldb"
	"agfun/auth/dto"
	"agfun/auth/entity"
	"agfun/crypto"
	"fmt"
	"github.com/satori/go.uuid"
	"time"
)

func (s *AuthSvc) CreateUser(req entity.User) (*entity.User, error) {
	users, _, e := mysqldb.GetUsers(req)
	if e != nil {
		return nil, e
	}
	if len(users) > 0 {
		return nil, fmt.Errorf("user is had")
	}
	req.CreateTime = time.Now()
	user, e := mysqldb.CreateUser(&req)
	if e != nil {
		return nil, e
	}
	return user, nil
}

func (s *AuthSvc) Login(req entity.User) (*entity.User, error) {
	users, _, e := mysqldb.GetUsers(req)
	if e != nil {
		return nil, e
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("user name or pwd is error")
	}
	accessToken := crypto.SHA1(uuid.Must(uuid.NewV4()).String())
	users[0].AccessToken = accessToken
	if e = mysqldb.UpdateUser(users[0]); e != nil {
		return nil, e
	}
	//update etcd
	e = s.Dynamic.PutAccessToken(accessToken, users[0].ID)
	if e != nil {
		return nil, e
	}

	return users[0], nil
}
func (s *AuthSvc) AddRole(vip *entity.UserRole, session string) error {
	var id string
	e := s.Dynamic.Get(session, &id)
	if e != nil {
		return e
	}
	if len(vip.UserID) == 0 {
		vip.UserID = id
	}
	e = mysqldb.AddRole(vip)
	if e != nil {
		return e
	}
	return nil
}
func (s *AuthSvc) GetUserRoles(vip entity.UserRole) ([]*entity.UserRole, error) {
	users, e := mysqldb.GetUserRoles(vip)
	if e != nil {
		return nil, e
	}
	return users, nil
}
func (s *AuthSvc) UpdateUserRole(vip *entity.UserRole, session string) error {
	if vip == nil {
		return fmt.Errorf("vip is nil")
	}
	src := entity.UserRole{
		ID:     "",
		UserID: vip.UserID,
		RoleID: vip.RoleID,
		Level:  0,
		Expire: time.Time{},
	}
	e := mysqldb.UpdateUserRole(vip, src)
	if e != nil {
		return e
	}
	return nil
}
func (s *AuthSvc) DelUserRole(userRole entity.UserRole, session string) error {
	e := mysqldb.DelUserRole(userRole)
	return e
}

func (s *AuthSvc) GetResources() ([]*dto.Resource, error) {
	return mysqldb.GetResources("0")
}
