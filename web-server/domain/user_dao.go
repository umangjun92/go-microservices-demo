package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/umangjun92/go-microservices-demo/web-server/utils"
)

type usersDao struct{}

type usersDaoInterface interface {
	GetUser(uint64) (*User, *utils.ApplicationError)
}

var (
	users = map[uint64]*User{
		123: {Id: 123, FirstName: "Umang", LastName: "Khan", Email: "uk@gmail.com"},
	}

	UsersDao usersDaoInterface = &usersDao{}
)

// func init() {
// 	UsersDao = &usersDao{}
// }

func (*usersDao) GetUser(userId uint64) (*User, *utils.ApplicationError) {
	fmt.Printf("calling DB")
	log.Printf("calling DB")
	user := users[userId]
	if user == nil {
		return nil, &utils.ApplicationError{Message: fmt.Sprintf("user %v not found", userId), StatusCode: http.StatusNotFound}
	}
	return user, nil
}
