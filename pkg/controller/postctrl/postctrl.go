package postctrl

import (
	"github.com/sylphritz/go-api-server/pkg/controller/crud"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"github.com/sylphritz/go-api-server/pkg/service/postservice"
)

func PostController() *crud.CrudCtrl[*schema.Post] {
	return crud.NewCrudController[*schema.Post](
		postservice.Name,
		schema.UserIdColumnName,
	)
}
