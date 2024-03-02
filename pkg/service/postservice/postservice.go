package postservice

import (
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"github.com/sylphritz/go-api-server/pkg/service"
)

const Name = "Post"

func NewService() *service.Service[schema.Post] {
	return service.NewService[schema.Post](Name)
}
