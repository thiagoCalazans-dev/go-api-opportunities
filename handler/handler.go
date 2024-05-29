package handler

import (
	"github.com/thiagoCalazans-dev/go-api-opportunities.git/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Init() {
	logger = config.GetLogger("handler")
	db = config.GetSqlite()
}
