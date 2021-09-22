package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umangjun92/go-microservices-demo/src/api/_utils/errors"
	"github.com/umangjun92/go-microservices-demo/src/api/domain/repos"
	"github.com/umangjun92/go-microservices-demo/src/api/services"
)

func CreateRepo(c *gin.Context) {
	var request repos.CreateRepoRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid JSON body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
