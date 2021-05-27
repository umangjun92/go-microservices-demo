package app

import (
	"github.com/umangjun92/go-microservices-demo/web-server/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)

}
