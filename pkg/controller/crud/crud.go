package crud

import (
	"github.com/gin-gonic/gin"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"github.com/sylphritz/go-api-server/pkg/service"
	"github.com/sylphritz/go-api-server/pkg/util"
)

type CrudController[T any] interface {
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type CrudCtrl[T schema.CommonEntity] struct {
	Name       string
	ForeignKey string
}

type Pagination struct {
	Page    int   `json:"page"`
	PerPage int   `json:"perPage"`
	Total   int64 `json:"total"`
}

type GetByIdRouteParams struct {
	ID uint `json:"id" uri:"id"`
}

type PaginatedRequest struct {
	Page    string `json:"page" form:"page"`
	PerPage string `json:"perPage" form:"perPage"`
}

type Response[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

type PaginatedResponse[T any] struct {
	Success    bool       `json:"success"`
	Data       []T        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type EmptyResponse[T any] struct {
	Success bool `json:"success"`
}

func GetPaginationQueryParams(c *gin.Context, r *PaginatedRequest) (int, int, bool) {
	page := util.StringToIntWithDefault(r.Page, service.DefaultPage)
	perPage := util.StringToIntWithDefault(r.PerPage, service.DefaultPerPage)

	return page, perPage, true
}

func NewCrudController[T schema.CommonEntity](name, foreignKey string) *CrudCtrl[T] {
	return &CrudCtrl[T]{name, foreignKey}
}
