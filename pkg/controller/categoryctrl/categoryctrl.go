package categoryctrl

import (
	"github.com/sylphritz/go-api-server/pkg/controller/crud"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"github.com/sylphritz/go-api-server/pkg/service/categoryservice"
)

func CategoryController() *crud.CrudCtrl[*schema.Category] {
	return crud.NewCrudController[*schema.Category](
		categoryservice.Name,
		schema.UserIdColumnName,
	)
}
