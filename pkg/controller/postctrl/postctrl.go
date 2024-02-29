package postctrl

import (
	"github.com/sylphritz/go-api-server/pkg/controller/crud"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"github.com/sylphritz/go-api-server/pkg/service/blogservice"
)

func PostController() *crud.CrudCtrl[*schema.Post] {
	return crud.NewController[*schema.Post](
		blogservice.Name,
		schema.UserIdColumnName,
	)
}
