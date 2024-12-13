package main

import (
	"log"
	"net/http"

	"github.com/brunosprado/validacao-cpf-cnpj.git/backend/database"
	"github.com/brunosprado/validacao-cpf-cnpj.git/backend/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Inicializa o banco de dados
	database.InitDatabase()

	// Cria o roteador
	r := mux.NewRouter()
	routes.RegisterCPFCNPJRoutes(r)

	// Inicia o servidor
	log.Println("Servidor iniciado!")
	log.Fatal(http.ListenAndServe(":8080", r))
}
