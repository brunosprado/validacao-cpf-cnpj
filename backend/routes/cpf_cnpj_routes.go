package routes

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterCPFCNPJRoutes(r *mux.Router, DB *gorm.DB) {
	h := handle{DB: DB}

	r.HandleFunc("/cpf-cnpj", h.GetAllRecords).Methods("GET")
	r.HandleFunc("/cpf-cnpj", h.CreateRecord).Methods("POST")
	r.HandleFunc("/cpf-cnpj/{id}", h.UpdateRecord).Methods("PUT")
	r.HandleFunc("/cpf-cnpj/{id}", h.DeleteRecord).Methods("DELETE")
	r.HandleFunc("/status", h.StatusHandler).Methods("GET")
}
