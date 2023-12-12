package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"github.com/sylphritz/go-api-server/pkg/db/service"
	"github.com/sylphritz/go-api-server/pkg/response"
	"github.com/sylphritz/go-api-server/pkg/util"
	"gorm.io/gorm"
)

type UserResponse struct {
	Data *schema.User `json:"data"`
}

func RegisterUser(c *gin.Context) {
	email, found := c.GetPostForm("email")
	if !found {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("%v is required", email),
		})
		return
	}

	password, found := c.GetPostForm("password")
	if !found {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("%v is required", password),
		})
		return
	}

	fmt.Println("Incoming user registration request:")
	fmt.Println("- email:", email)
	fmt.Println("- password:", password)

	hashedPassword, err := util.GenerateHashedPassword(password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error generating password hash",
		})
		return
	}

	user := schema.User{
		Email:    email,
		Password: hashedPassword,
	}

	err = service.CreateUser(&user)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.AbortWithStatusJSON(http.StatusConflict, response.ErrorResponse{
				Status:  http.StatusConflict,
				Message: "Email already exists",
			})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, UserResponse{&user})
}
