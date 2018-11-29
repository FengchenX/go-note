package service

import (
	"agfun/agfun-service/crypto"
	"agfun/auth/dbcentral/mysqldb"
	"agfun/auth/entity"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (s *AuthSvc) CreateUser(req entity.User) (*entity.User, error) {
	users, _, e := mysqldb.GetUsers(req)
	if e != nil {
		return nil, e
	}
	if len(users) > 0 {
		return nil, fmt.Errorf("user is had")
	}
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
func (s *AuthSvc) AddVip(vip *entity.VipUser, session string) error {
	var id int
	e := s.Dynamic.Get(session, &id)
	if e != nil {
		return e
	}
	if vip.UserID == 0 {
		vip.UserID = uint(id)
	}
	e = mysqldb.AddVip(vip)
	if e != nil {
		return e
	}
	return nil
}
func (s *AuthSvc) GetVips(vip entity.VipUser) ([]*entity.VipUser, error) {
	users, e := mysqldb.GetVips(vip)
	if e != nil {
		return nil, e
	}
	return users, nil
}
func (s *AuthSvc) UpdateVip(vip *entity.VipUser, session string) error {
	if vip == nil {
		return fmt.Errorf("vip is nil")
	}
	src := entity.VipUser{
		Model: gorm.Model{
			ID: vip.ID,
		},
	}
	e := mysqldb.UpdateVip(vip, src)
	if e != nil {
		return e
	}
	return nil
}
func (s *AuthSvc) DelVip(id int, session string) error {
	user := entity.VipUser{}
	user.ID = uint(id)
	e := mysqldb.DelVip(user)
	return e
}
