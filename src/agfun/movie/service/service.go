package service

import (
	"service"
)

type MovieSvc struct {
	service.Svc
}

func NewMovieSvc() *MovieSvc {
	return &MovieSvc{Svc: *service.DefaultSvc()}
}
