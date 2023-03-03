package service

import "nagato/internal/data"

type Service interface {
	Conf() ConfSrv
}

type service struct {
	store *data.Store
}

func (s service) Conf() ConfSrv {
	return newConfService(s.store)
}

func NewService() Service {
	return &service{store: data.GetStore()}
}
