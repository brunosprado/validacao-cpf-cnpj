package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	dsn := "host=localhost user=postgres password=dinossauro dbname=cpf_cnpj port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}
}
