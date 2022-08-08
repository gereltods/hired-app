package services

import (
	"fasthttp/domains/db"
)

var UsersService usersServiceInterface = &usersService{}

type usersService struct{}

type usersServiceInterface interface {
	GetAll() []byte
	GetAllWithoutCache() []byte
}

func (s *usersService) GetAll() []byte {
	return db.LoadAllUser()
}

func (s *usersService) GetAllWithoutCache() []byte {
	return db.LoadAllUserNoCache()
}
