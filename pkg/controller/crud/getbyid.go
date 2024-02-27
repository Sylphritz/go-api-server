package crud

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controller"
	"github.com/sylphritz/go-api-server/pkg/response"
	"github.com/sylphritz/go-api-server/pkg/service"
	"github.com/sylphritz/go-api-server/pkg/session"
)

func (crud CrudCtrl[T]) GetById(c *gin.Context) {
	serv := service.NewService[T](crud.Name)

	var params GetByIdRouteParams

	ok := controller.GetRouteParams[GetByIdRouteParams](c, &params)
	if !ok {
		return
	}

	userInfo := c.MustGet(session.UserKey).(*session.UserSessionInfo)
	foreignKey := service.ForeignKeyQuery{Key: crud.ForeignKey, Value: userInfo.ID}

	item, err := serv.FindById(params.ID, &foreignKey)
	if err != nil {
		response.RespondInternalErrorWithMessage(c, err)
		return
	}

	c.JSON(http.StatusOK, Response[T]{
		true,
		*item,
	})
}
