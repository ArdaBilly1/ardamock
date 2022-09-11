package config

import (
	"ardamock/src/model"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		model.Endpoint{},
		model.Request{},
		model.Response{},
	)
}
