package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const InternalServerErrorMessage = "Internal server error."

func NewErrorResponse(code int, message string) ErrorResponse {
	return ErrorResponse{code, message}
}

func NewInternalServerError() ErrorResponse {
	return NewErrorResponse(http.StatusInternalServerError, InternalServerErrorMessage)
}

func NewBadRequestError(message string) ErrorResponse {
	return NewErrorResponse(http.StatusBadRequest, message)
}

func NewConflictError(message string) ErrorResponse {
	return NewErrorResponse(http.StatusConflict, message)
}

func RespondWithError(c *gin.Context, err ErrorResponse) {
	c.AbortWithStatusJSON(err.Code, err)
}

func RespondInternalErrorWithMessage(c *gin.Context, err error) {
	RespondWithError(c, NewErrorResponse(
		http.StatusInternalServerError,
		err.Error(),
	))
}
