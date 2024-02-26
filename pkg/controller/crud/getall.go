package crud

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/controller"
	"github.com/sylphritz/go-api-server/pkg/response"
	"github.com/sylphritz/go-api-server/pkg/service"
	"github.com/sylphritz/go-api-server/pkg/session"
)

func (crud CrudCtrl[T]) GetAll(c *gin.Context) {
	serv := service.NewService[T](crud.Name)

	var request PaginatedRequest

	ok := controller.GetRequiredQueryParams[PaginatedRequest](c, &request)
	if !ok {
		log.Println("No query params")
		log.Println(request)
		return
	}

	page, perPage, ok := GetPaginationQueryParams(c, &request)
	if !ok {
		return
	}

	userInfo := c.MustGet(session.UserKey).(*session.UserSessionInfo)
	foreignKey := service.ForeignKeyQuery{Key: crud.ForeignKey, Value: userInfo.ID}

	items, err := serv.FindAll(page, perPage, &foreignKey)
	if err != nil {
		response.RespondInternalErrorWithMessage(c, err)
	}

	count, err := serv.Count(&foreignKey)
	if err != nil {
		response.RespondInternalErrorWithMessage(c, err)
	}

	c.JSON(http.StatusOK, PaginatedResponse[T]{
		true,
		*items,
		Pagination{page, perPage, count},
	})
}
