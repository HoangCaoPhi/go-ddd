package shared

import (
	"phihc116/go-task/backend/webapi/options"

	"gorm.io/gorm"
)

var (
	Config    options.Config
	DbContext *gorm.DB
)
