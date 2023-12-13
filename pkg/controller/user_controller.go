package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"github.com/sylphritz/go-api-server/pkg/db/service"
	"github.com/sylphritz/go-api-server/pkg/response"
	"github.com/sylphritz/go-api-server/pkg/util"
)

type UserResponse struct {
	Data *schema.User `json:"data"`
}

func abortRequiredFieldError(c *gin.Context, fieldName string) {
	response.RespondWithError(c, response.NewBadRequestError(fmt.Sprintf("%v is required", fieldName)))
}

func getRequiredFields(c *gin.Context, fields []string) (map[string]string, bool) {
	data := map[string]string{}

	for _, field := range fields {
		value, found := c.GetPostForm(field)
		if !found {
			abortRequiredFieldError(c, field)
			return data, false
		}

		data[field] = value
	}

	return data, true
}

func isDuplicateEmail(err error) bool {
	return strings.Contains(err.Error(), "duplicate key")
}

func RegisterUser(c *gin.Context) {
	data, ok := getRequiredFields(c, []string{"email", "password"})
	if !ok {
		return
	}

	hashedPassword, err := util.GenerateHashedPassword(data["password"])
	if err != nil {
		response.RespondWithError(c, response.NewInternalServerError())
		return
	}

	user := schema.User{
		Email:    data["email"],
		Password: hashedPassword,
	}

	err = service.CreateUser(&user)
	if err != nil {
		if isDuplicateEmail(err) {
			response.RespondWithError(c, response.NewConflictError("Email already exists."))
			return
		} else {
			response.RespondWithError(c, response.NewInternalServerError())
			return
		}
	}

	c.JSON(http.StatusOK, UserResponse{&user})
}
