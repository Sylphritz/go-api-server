package blogservice

import (
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"github.com/sylphritz/go-api-server/pkg/service"
)

const Name = "Category"

func NewCategoryService() *service.Service[schema.Category] {
	return service.NewService[schema.Category](Name)
}
