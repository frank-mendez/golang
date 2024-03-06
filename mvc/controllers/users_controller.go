package controllers

import (
	"github.com/gin-gonic/gin"
	"golang/mvc/services"
	"golang/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "not_found",
		}
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUsers(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}
