package utils

import "github.com/gin-gonic/gin"

func Respond(c *gin.Context, statusCode int, data interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(statusCode, data)
		return
	}

	c.JSON(statusCode, data)
}

func RespondError(c *gin.Context, err *ApplicationError) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(err.StatusCode, err)
		return
	}

	c.JSON(err.StatusCode, err)
}
