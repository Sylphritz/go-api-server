package schema

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Task string
}
