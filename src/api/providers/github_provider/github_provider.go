package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/umangjun92/go-microservices-demo/src/api/_utils"
	restclient "github.com/umangjun92/go-microservices-demo/src/api/clients/rest_client"
	"github.com/umangjun92/go-microservices-demo/src/api/domain/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	response, err := restclient.Post(_utils.URL_CREATE_REPO, request, headers)

	if err != nil {
		log.Printf("err trying to create new repo: %s", err)
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	resBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(resBytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "Invalid JSON response body"}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse

	if err := json.Unmarshal(resBytes, &result); err != nil {
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "Invalid JSON response body"}
	}

	return &result, nil
}
