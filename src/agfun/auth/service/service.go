// Package classification Auth API.
//
// The purpose of this service is to provide an application
// that is using plain go code to define an API
//
// Schemes: http, https
//   Host: localhost:8080
//   Version: 0.0.1
//
// swagger:meta
package service

import (
	"agfun/agfun-service/service"
	"agfun/auth/dbcentral/etcd"
)

type AuthSvc struct {
	*service.Svc
	Dynamic *etcd.Client
}

func NewAuthSvc() *AuthSvc {
	return &AuthSvc{}
}

var stdSvc *AuthSvc

func initStdSvc() {
	if stdSvc != nil {
		return
	}
	stdSvc = NewAuthSvc()

	stdSvc.Svc = service.GetDefaultSvc()

	stdSvc.Dynamic = etcd.GetDefaultCli()
}
func GetDefaultSvc() *AuthSvc {
	if stdSvc == nil {
		initStdSvc()
	}
	return stdSvc
}
