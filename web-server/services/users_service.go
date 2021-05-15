package services

import (
	"github.com/umangjun92/go-microservices-demo/web-server/domain"
	"github.com/umangjun92/go-microservices-demo/web-server/utils"
)

func GetUser(userId uint64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
