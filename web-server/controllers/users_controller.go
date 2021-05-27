package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/umangjun92/go-microservices-demo/web-server/services"
	"github.com/umangjun92/go-microservices-demo/web-server/utils"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		userErr := utils.ApplicationError{Message: "user_id must be a number", StatusCode: http.StatusBadRequest}
		utils.Respond(c, int(userErr.StatusCode), userErr)
		// jsonVal, _ := json.Marshal(userErr)
		// resp.WriteHeader(int(userErr.StatusCode))
		// resp.Write(jsonVal)
		return
	}

	user, _err := services.UsersService.GetUser(uint64(userId))

	if _err != nil {
		// jsonVal, _ := json.Marshal(_err)
		// resp.WriteHeader(int(_err.StatusCode))
		// resp.Write(jsonVal)
		utils.Respond(c, int(_err.StatusCode), _err)
		return
	}
	utils.Respond(c, http.StatusOK, user)
}
