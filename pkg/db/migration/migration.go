package migration

import (
	"github.com/sylphritz/go-api-server/pkg/db/schema"
	"gorm.io/gorm"
)

func MigrateSchemas(db *gorm.DB) {
	db.AutoMigrate(
		&schema.Todo{},
		&schema.User{},
	)
}
