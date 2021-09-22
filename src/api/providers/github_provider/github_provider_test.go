package github_provider

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/umangjun92/go-microservices-demo/src/api/_utils"
	restclient "github.com/umangjun92/go-microservices-demo/src/api/clients/rest_client"
	"github.com/umangjun92/go-microservices-demo/src/api/domain/github"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("acb123")
	assert.EqualValues(t, "token acb123", header)
}

func TestCreateRepoErrorRestClient(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{Url: _utils.URL_CREATE_REPO, HttpMethod: http.MethodPost, Err: errors.New("invalid rest client response")})
	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid rest client response", err.Message)

	restclient.StopMockups()
}

// func TestCreateRepoInvalidResponseBody(t *testing.T) {
// 	restclient.StartMockups()
// 	restclient.FlushMocks()

// 	invalidCloser, _ := os.Open("-asf3")
// 	restclient.AddMockup(restclient.Mock{Url: _utils.URL_CREATE_REPO, HttpMethod: http.MethodPost, Response: &http.Response{StatusCode: http.StatusCreated, Body: invalidCloser}})

// 	response, err := CreateRepo("", github.CreateRepoRequest{})

// 	assert.Nil(t, response)
// 	assert.NotNil(t, err)
// 	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
// 	assert.EqualValues(t, "Invalid JSON response body", err.Message)

// 	restclient.StopMockups()
// }

func TestCreateRepoInvalidErrInterface(t *testing.T) {
	restclient.StartMockups()
	restclient.FlushMocks()

	restclient.AddMockup(restclient.Mock{
		Url:        _utils.URL_CREATE_REPO,
		HttpMethod: http.MethodPost,
		Response:   &http.Response{StatusCode: http.StatusUnauthorized, Body: ioutil.NopCloser(strings.NewReader(`{"message": 1}`))},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	j, _ := json.Marshal(response)
	fmt.Printf("response %s \n", j)
	e, _ := json.Marshal(err)
	fmt.Printf("error %s \n", e)
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid JSON response body", err.Message)

}
