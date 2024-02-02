package userctrl

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controller"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"github.com/sylphritz/go-api-server/pkg/response"
	"github.com/sylphritz/go-api-server/pkg/service/userservice"
	"github.com/sylphritz/go-api-server/pkg/util"
)

type UserResponse struct {
	Data *schema.User `json:"data"`
}

func isDuplicateEmail(err error) bool {
	return strings.Contains(err.Error(), "duplicate key")
}

func RegisterUser(c *gin.Context) {
	data, ok := controller.GetRequiredFormFields(c, []string{"email", "password"})
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

	err = userservice.CreateUser(&user)
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
