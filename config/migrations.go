package config

import (
	"sigma/api/models"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

func MigrateTables(db *gorm.DB) {
	db.AutoMigrate(&models.BusinessUser{})
}
