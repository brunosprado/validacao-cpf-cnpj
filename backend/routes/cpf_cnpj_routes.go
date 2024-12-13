package routes

import (
	"github.com/brunosprado/validacao-cpf-cnpj.git/backend/handlers"
	"github.com/gorilla/mux"
)

func RegisterCPFCNPJRoutes(r *mux.Router) {
	r.HandleFunc("/api/cpfs", handlers.GetAllRecords).Methods("GET")
	r.HandleFunc("/api/cpfs", handlers.CreateRecord).Methods("POST")
	r.HandleFunc("/api/cpfs/{id}", handlers.UpdateRecord).Methods("PUT")
	r.HandleFunc("/api/cpfs/{id}", handlers.DeleteRecord).Methods("DELETE")
	r.HandleFunc("/status", handlers.StatusHandler).Methods("GET")
}
