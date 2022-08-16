package services

import (
	"fasthttp/domains/db"
)

var UsersService usersServiceInterface = &usersService{}

type usersService struct{}

type usersServiceInterface interface {
	GetAll() *db.Dbs
	GetAllWithoutCache() []byte
}

func (s *usersService) GetAll() *db.Dbs {
	result := &db.Dbs{}
	result.LoadAllUser()
	return result
}

func (s *usersService) GetAllWithoutCache() []byte {
	return db.LoadAllUserNoCache()
}
