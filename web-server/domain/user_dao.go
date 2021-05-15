package domain

import (
	"fmt"
	"net/http"

	"github.com/umangjun92/go-microservices-demo/web-server/utils"
)

var (
	users = map[uint64]*User{
		123: {Id: 123, FirstName: "Umang", LastName: "Khan", Email: "uk@gmail.com"},
	}
)

func GetUser(userId uint64) (*User, *utils.ApplicationError) {
	user := users[userId]
	if user == nil {
		return nil, &utils.ApplicationError{Message: fmt.Sprintf("user %v not found", userId), StatusCode: http.StatusNotFound}
	}
	return user, nil
}
