package repository

import (
	"github.com/ADStefano/NimbusAPI/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conect() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
