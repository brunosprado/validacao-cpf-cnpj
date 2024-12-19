package models

import (
	"gorm.io/gorm"
)

type CPF_CNPJ struct {
	gorm.Model
	ID        uint   `json:"id"`
	Number    string `gorm:"not null;unique" json:"number"`
	IsBlocked bool   `json:"is_blocked"`
}
