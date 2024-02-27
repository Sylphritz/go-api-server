package crud

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controller"
	"github.com/sylphritz/go-api-server/pkg/response"
	"github.com/sylphritz/go-api-server/pkg/service"
	"github.com/sylphritz/go-api-server/pkg/session"
)

func (crud CrudCtrl[T]) Create(c *gin.Context) {
	serv := service.NewService[T](crud.Name)

	var entity T

	ok := controller.GetRequiredBodyFields[T](c, &entity)
	if !ok {
		return
	}

	userInfo := c.MustGet(session.UserKey).(*session.UserSessionInfo)
	entity.SetUserID(userInfo.ID)

	err := serv.Create(&entity)
	if err != nil {
		response.RespondInternalErrorWithMessage(c, err)
		return
	}

	c.JSON(http.StatusOK, Response[T]{
		true,
		entity,
	})
}
