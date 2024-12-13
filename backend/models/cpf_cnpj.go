package models

type CPF_CNPJ struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Number    string `gorm:"unique;not null" json:"number"`
	IsBlocked bool   `json:"isBlocked"`
}
