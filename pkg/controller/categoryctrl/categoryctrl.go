package categoryctrl

import (
	"github.com/sylphritz/go-api-server/pkg/controller/crud"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"github.com/sylphritz/go-api-server/pkg/service/blogservice"
)

func CategoryController() *crud.CrudCtrl[*schema.Category] {
	return crud.NewController[*schema.Category](
		blogservice.Name,
		schema.UserIdColumnName,
	)
}
