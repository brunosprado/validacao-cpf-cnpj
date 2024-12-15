package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=dinossauro dbname=postgres port=5433 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
