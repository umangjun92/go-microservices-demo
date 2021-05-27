package services

import (
	"github.com/umangjun92/go-microservices-demo/web-server/domain"
	"github.com/umangjun92/go-microservices-demo/web-server/utils"
)

type usersService struct{}

var UsersService usersService

func (*usersService) GetUser(userId uint64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UsersDao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
