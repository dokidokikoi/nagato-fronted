package service

import "nagato/internal/data"

type Service interface {
	Conf() ConfSrv
	File() FileSrv
	Blank() BlankSrv
}

type service struct {
	store *data.Store
}

func (s service) Conf() ConfSrv {
	return newConfService(s.store)
}

func (s service) File() FileSrv {
	return newFileService(s.store)
}

func (s service) Blank() BlankSrv {
	return newBlankService(s.store)
}

func NewService() Service {
	return &service{store: data.GetStore()}
}

var Authorization = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZW1haWwiOiIxMjNAZXhhbXBsZS5jb20iLCJyb2xlIjoiIiwiZXhwIjoxNzE0NDU3NjE2LCJpYXQiOjE2ODMwMDgwMTYsImlzcyI6ImhhcnVrYXplIiwibmJmIjoxNjgzMDA4MDE2fQ.c1VZqJEtNSKeekzfi1WSd113UobIhtdvvlKn7KYG4kg"
