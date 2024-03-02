package pagectrl

import (
	"github.com/sylphritz/go-api-server/pkg/controller/crud"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"github.com/sylphritz/go-api-server/pkg/service/pageservice"
)

func PageController() *crud.CrudCtrl[*schema.Page] {
	return crud.NewCrudController[*schema.Page](
		pageservice.Name,
		schema.UserIdColumnName,
	)
}
