package app

import (
	"github.com/umangjun92/go-microservices-demo/src/api/controllers/polo"
	"github.com/umangjun92/go-microservices-demo/src/api/controllers/repos"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repos", repos.CreateRepo)
}
