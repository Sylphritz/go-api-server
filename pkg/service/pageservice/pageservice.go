package pageservice

import (
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"github.com/sylphritz/go-api-server/pkg/service"
)

const Name = "Page"

func NewService(foreignKey *service.ForeignKeyQuery) *service.Service[schema.Page] {
	return service.NewService[schema.Page](Name)
}
