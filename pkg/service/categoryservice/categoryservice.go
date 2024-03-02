package categoryservice

import (
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"github.com/sylphritz/go-api-server/pkg/service"
)

const Name = "Category"

func NewService() *service.Service[schema.Category] {
	return service.NewService[schema.Category](Name)
}
