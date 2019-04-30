package service

import (
	"service"
)

type MovieSvc struct {
	Svc *service.Svc
}

func NewMovieSvc() *MovieSvc {
	return &MovieSvc{Svc: service.DefaultSvc()}
}

//var stdSvc *MovieSvc
//
//func initStdSvc() {
//	if stdSvc != nil {
//		return
//	}
//	stdSvc = NewMovieSvc()
//	stdSvc.Svc = service.GetDefaultSvc()
//}
//func GetDefaultSvc() *MovieSvc {
//	if stdSvc == nil {
//		initStdSvc()
//	}
//	return stdSvc
//}
