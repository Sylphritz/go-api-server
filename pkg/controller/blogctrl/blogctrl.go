package blogctrl

import (
	"github.com/sylphritz/go-api-server/pkg/controller/crud"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"github.com/sylphritz/go-api-server/pkg/service/blogservice"
)

func BlogControllerByUser() *crud.CrudCtrl[schema.Blog] {
	return crud.NewController[schema.Blog](
		blogservice.Name,
		"user_id",
	)
}
