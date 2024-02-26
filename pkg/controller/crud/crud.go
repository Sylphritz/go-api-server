package crud

import (
	"github.com/gin-gonic/gin"
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

type CrudCtrl[T any] struct {
	Name       string
	ForeignKey string
}

type Pagination struct {
	Page    int   `json:"page"`
	PerPage int   `json:"perPage"`
	Total   int64 `json:"total"`
}

type PaginatedRequest struct {
	Page    string `json:"page" form:"page"`
	PerPage string `json:"perPage" form:"perPage"`
}

type PaginatedResponse[T any] struct {
	Success    bool       `json:"success"`
	Data       []T        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func GetPaginationQueryParams(c *gin.Context, r *PaginatedRequest) (int, int, bool) {
	page := util.StringToIntWithDefault(r.Page, service.DefaultPage)
	perPage := util.StringToIntWithDefault(r.PerPage, service.DefaultPerPage)

	return page, perPage, true
}

func NewController[T any](name, foreignKey string) *CrudCtrl[T] {
	return &CrudCtrl[T]{name, foreignKey}
}
