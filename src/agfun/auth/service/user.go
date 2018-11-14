package service

import (
	"agfun/agfun-service/util"
	"agfun/auth/dbcentral/mysqldb"
	"agfun/auth/entity"
	"fmt"
	"github.com/satori/go.uuid"
)

func CreateUser(req entity.User) (*entity.User, error) {
	users, e := mysqldb.GetUsers(req)
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

func Login(req entity.User) (*entity.User, error) {
	users, e := mysqldb.GetUsers(req)
	if e != nil {
		return nil, e
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("user name or pwd is error")
	}
	accessToken := util.SHA1(uuid.Must(uuid.NewV4()).String())
	users[0].AccessToken = accessToken
	if e = mysqldb.UpdateUser(users[0]); e != nil {
		return nil, e
	}
	//update redis

	return users[0], nil
}
