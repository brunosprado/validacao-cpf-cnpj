package main

import (
	"log"
	"net/http"

	"github.com/brunosprado/validacao-cpf-cnpj.git/backend/database"
	"github.com/brunosprado/validacao-cpf-cnpj.git/backend/models"
	"github.com/brunosprado/validacao-cpf-cnpj.git/backend/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Inicializa o banco de dados e migrate
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.CPF_CNPJ{})

	// Cria o roteador
	r := mux.NewRouter()
	routes.RegisterCPFCNPJRoutes(r, db)
	r.Use(routes.CorsMiddleware)

	// Inicia o servidor
	log.Println("Servidor iniciado!")
	log.Fatal(http.ListenAndServe(":8080", r))
}
