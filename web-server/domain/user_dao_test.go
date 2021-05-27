package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := UsersDao.GetUser(0)

	assert.Nil(t, user, "No user expected")

	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)

}

// func TestGetUserWithNoError(t *testing.T) {
// 	GetUser(123)
// }
