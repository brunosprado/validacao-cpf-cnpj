package routes

import (
	"github.com/brunosprado/validacao-cpf-cnpj.git/backend/handlers"
	"github.com/gorilla/mux"
)

func RegisterCPFCNPJRoutes(r *mux.Router) {
	r.HandleFunc("/cpf-cnpj", handlers.GetAllRecords).Methods("GET")
	r.HandleFunc("/cpf-cnpj", handlers.CreateRecord).Methods("POST")
	r.HandleFunc("/cpf-cnpj/{id}", handlers.UpdateRecord).Methods("PUT")
	r.HandleFunc("/cpf-cnpj/{id}", handlers.DeleteRecord).Methods("DELETE")
	r.HandleFunc("/status", handlers.StatusHandler).Methods("GET")
}
