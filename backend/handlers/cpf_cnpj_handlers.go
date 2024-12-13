package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/brunosprado/validacao-cpf-cnpj.git/backend/models"
	"github.com/brunosprado/validacao-cpf-cnpj.git/backend/services"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB
var QueriesCount int

func GetAllRecords(w http.ResponseWriter, r *http.Request) {
	QueriesCount++
	var records []models.CPF_CNPJ
	if err := DB.Find(&records).Error; err != nil {
		http.Error(w, "Erro ao buscar registros", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	QueriesCount++
	var record models.CPF_CNPJ
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		http.Error(w, "Payload inválido", http.StatusBadRequest)
		return
	}

	if !services.ValidateCPF_CNPJ(record.Number) {
		http.Error(w, "CPF/CNPJ inválido", http.StatusBadRequest)
		return
	}

	if err := DB.Create(&record).Error; err != nil {
		http.Error(w, "Erro ao criar registro", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func UpdateRecord(w http.ResponseWriter, r *http.Request) {
	QueriesCount++
	params := mux.Vars(r)
	var record models.CPF_CNPJ

	if err := DB.First(&record, params["id"]).Error; err != nil {
		http.Error(w, "Registro não encontrado", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		http.Error(w, "Payload inválido", http.StatusBadRequest)
		return
	}

	if !services.ValidateCPF_CNPJ(record.Number) {
		http.Error(w, "CPF/CNPJ inválido", http.StatusBadRequest)
		return
	}

	DB.Save(&record)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	QueriesCount++
	params := mux.Vars(r)
	var record models.CPF_CNPJ

	if err := DB.Delete(&record, params["id"]).Error; err != nil {
		http.Error(w, "Erro ao deletar registro", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	uptime := map[string]interface{}{
		"uptime":       "running",
		"queriesCount": QueriesCount,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uptime)
}
