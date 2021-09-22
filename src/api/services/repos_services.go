package services

import (
	"strings"

	"github.com/umangjun92/go-microservices-demo/src/api/_utils/errors"
	"github.com/umangjun92/go-microservices-demo/src/api/config"
	"github.com/umangjun92/go-microservices-demo/src/api/domain/github"
	"github.com/umangjun92/go-microservices-demo/src/api/domain/repos"
	"github.com/umangjun92/go-microservices-demo/src/api/providers/github_provider"
)

type reposService struct {
}

type reposServiceInterface interface {
	CreateRepo(request repos.CreateRepoRequest) (*repos.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(input repos.CreateRepoRequest) (*repos.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("Invalid repository name")
	}
	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: "",
		Private:     false,
	}
	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)

	if err != nil {
		apiErr := errors.NewApiError(err.StatusCode, err.Message)
		return nil, apiErr
	}

	result := repos.CreateRepoResponse{Id: response.Id, Name: response.Name, Owner: response.Owner.Login}
	return &result, nil
}
