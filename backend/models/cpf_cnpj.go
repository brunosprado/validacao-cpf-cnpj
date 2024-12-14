package models

import (
	"time"

	"gorm.io/gorm"
)

type CPF_CNPJ struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Number    string    `gorm:"type:varchar(14);unique;not null" json:"number"`
	IsBlocked bool      `gorm:"default:false" json:"is_blocked"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
