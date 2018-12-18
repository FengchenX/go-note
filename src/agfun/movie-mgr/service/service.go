// Package classification movie API.
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

import "agfun/agfun-service/service"

type MovieSvc struct {
	*service.Svc
}

func NewMovieSvc() *MovieSvc {
	return &MovieSvc{}
}

var stdSvc *MovieSvc

func initStdSvc() {
	if stdSvc != nil {
		return
	}
	stdSvc = NewMovieSvc()
	stdSvc.Svc = service.GetDefaultSvc()
}
func GetDefaultSvc() *MovieSvc {
	if stdSvc == nil {
		initStdSvc()
	}
	return stdSvc
}
