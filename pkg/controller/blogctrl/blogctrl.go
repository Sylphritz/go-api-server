package blogctrl

import (
	"github.com/sylphritz/go-api-server/pkg/controller/crud"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"github.com/sylphritz/go-api-server/pkg/service/blogservice"
)

func BlogController() *crud.CrudCtrl[*schema.Blog] {
	return crud.NewCrudController[*schema.Blog](
		blogservice.Name,
		schema.UserIdColumnName,
	)
}
