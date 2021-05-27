package services_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/umangjun92/go-microservices-demo/web-server/domain"
	"github.com/umangjun92/go-microservices-demo/web-server/services"
	"github.com/umangjun92/go-microservices-demo/web-server/utils"
)

type usersDaoMock struct {
}

var (
	// _usersDaoMock usersDaoMock
	getUserFunc func(userId uint64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UsersDao = &usersDaoMock{}
}

func (*usersDaoMock) GetUser(userId uint64) (*domain.User, *utils.ApplicationError) {
	return getUserFunc(userId)
}

func TestUsersServicesNotFound(t *testing.T) {
	getUserFunc = func(user uint64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{Message: "Not found", StatusCode: http.StatusNotFound}
	}
	user, err := services.UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestUsersServicesNoError(t *testing.T) {
	getUserFunc = func(user uint64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{Id: 123}, nil
	}
	user, err := services.UsersService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
}
