package blogservice

import (
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"github.com/sylphritz/go-api-server/pkg/service"
)

const Name = "Blog"

func NewService(foreignKey *service.ForeignKeyQuery) *service.Service[schema.Blog] {
	return service.NewService[schema.Blog](Name)
}
