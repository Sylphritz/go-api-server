package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/response"
)

func AbortRequiredFieldError(c *gin.Context, fieldName string) {
	response.RespondWithError(c, response.NewBadRequestError(fmt.Sprintf("%v is required", fieldName)))
}

func GetRequiredFormFields(c *gin.Context, fields []string) (map[string]string, bool) {
	data := map[string]string{}

	for _, field := range fields {
		value, found := c.GetPostForm(field)
		if !found {
			AbortRequiredFieldError(c, field)
			return data, false
		}

		data[field] = value
	}

	return data, true
}

func GetRequiredBodyFields[S any](c *gin.Context, b *S) bool {
	if err := c.BindJSON(&b); err != nil {
		response.RespondWithError(c, response.NewBadRequestError(err.Error()))
		return false
	}

	return true
}
