package handler

import (
	"github.com/thiago-kersting/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("Handler")
	db = config.GetSQLite()
}
