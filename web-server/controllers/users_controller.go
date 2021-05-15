package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/umangjun92/go-microservices-demo/web-server/services"
	"github.com/umangjun92/go-microservices-demo/web-server/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		userErr := utils.ApplicationError{Message: "user_id must be a number", StatusCode: http.StatusBadRequest}
		jsonVal, _ := json.Marshal(userErr)
		resp.WriteHeader(int(userErr.StatusCode))
		resp.Write(jsonVal)
		return
	}

	user, _err := services.GetUser(uint64(userId))
	if _err != nil {
		jsonVal, _ := json.Marshal(_err)
		resp.WriteHeader(int(_err.StatusCode))
		resp.Write(jsonVal)
		return
	}
	jsonVal, _ := json.Marshal(user)
	resp.Write(jsonVal)
}
