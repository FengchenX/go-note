package service

import "agfun/service"

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
