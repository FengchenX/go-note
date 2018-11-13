package service

import (
	"agfun/auth/dbcentral/mysqldb"
	"agfun/auth/entity"
	"fmt"
)

func CreateUser(req entity.User) (*entity.User, error) {
	users, e := mysqldb.GetUsers(req)
	if e != nil {
		return nil, e
	}
	if len(users) > 0 {
		return nil, fmt.Errorf("user is had")
	}
	user, e := mysqldb.CreateUser(req)
	if e != nil {
		return nil, e
	}
	return user, nil
}
