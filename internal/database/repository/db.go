package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conect(cstr string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(cstr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
