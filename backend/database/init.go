package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() (*gorm.DB, error) {
	dsn := "host=db user=postgres password=dinossauro dbname=postgres port=5432 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
