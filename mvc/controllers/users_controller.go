package controllers

import (
	"encoding/json"
	"golang/mvc/services"
	"golang/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUsers(userId)
	if apiErr != nil {
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(apiErr.Message))
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
