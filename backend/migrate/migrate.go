package main

import (
	"github.com/brunosprado/validacao-cpf-cnpj.git/backend/database"
	"github.com/brunosprado/validacao-cpf-cnpj.git/backend/models"
)

func main() {
	database.InitDatabase()
	database.DB.AutoMigrate(&models.CPF_CNPJ{})
}
