package routes

import (
	"encoding/json"
	"net/http"

	"github.com/brunosprado/validacao-cpf-cnpj.git/backend/models"
	"github.com/brunosprado/validacao-cpf-cnpj.git/backend/services"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type handle struct {
	DB           *gorm.DB
	QueriesCount int
}

func (h *handle) GetAllRecords(w http.ResponseWriter, r *http.Request) {
	h.QueriesCount++
	var records []models.CPF_CNPJ
	if err := h.DB.Find(&records).Error; err != nil {
		http.Error(w, "Erro ao buscar registros", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}

func (h *handle) CreateRecord(w http.ResponseWriter, r *http.Request) {
	h.QueriesCount++
	var record models.CPF_CNPJ
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		http.Error(w, "Payload inválido", http.StatusBadRequest)
		return
	}

	if !services.ValidateCPF_CNPJ(record.Number) {
		http.Error(w, "CPF/CNPJ inválido", http.StatusBadRequest)
		return
	}

	if err := h.DB.Create(&record).Error; err != nil {
		http.Error(w, "Erro ao criar registro", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func (h *handle) UpdateRecord(w http.ResponseWriter, r *http.Request) {
	h.QueriesCount++
	params := mux.Vars(r)
	var record models.CPF_CNPJ

	if err := h.DB.First(&record, params["id"]).Error; err != nil {
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

	h.DB.Save(&record)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func (h *handle) DeleteRecord(w http.ResponseWriter, r *http.Request) {
	h.QueriesCount++
	params := mux.Vars(r)
	var record models.CPF_CNPJ

	if err := h.DB.Delete(&record, params["id"]).Error; err != nil {
		http.Error(w, "Erro ao deletar registro", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handle) StatusHandler(w http.ResponseWriter, r *http.Request) {
	uptime := map[string]interface{}{
		"uptime":       "running",
		"queriesCount": h.QueriesCount,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uptime)
}
