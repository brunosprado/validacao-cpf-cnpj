package services

import (
	"regexp"
	"strconv"
	"strings"
)

// Valida CPF ou CNPJ
func ValidateCPF_CNPJ(number string) bool {
	number = strings.ReplaceAll(number, ".", "")
	number = strings.ReplaceAll(number, "-", "")
	number = strings.ReplaceAll(number, "/", "")

	if len(number) == 11 {
		return validateCPF(number)
	} else if len(number) == 14 {
		return validateCNPJ(number)
	}
	return false
}

func validateCPF(cpf string) bool {
	if matched, _ := regexp.MatchString(`^[0-9]{11}$`, cpf); !matched {
		return false
	}

	// Ignora CPFs com todos os dígitos iguais (e.g., 11111111111)
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			break
		}
		if i == len(cpf)-1 {
			return false
		}
	}

	// Calcula o primeiro dígito verificador
	sum := 0
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (10 - i)
	}
	remainder := (sum * 10) % 11
	if remainder == 10 {
		remainder = 0
	}
	if remainder != int(cpf[9]-'0') {
		return false
	}

	// Calcula o segundo dígito verificador
	sum = 0
	for i := 0; i < 10; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (11 - i)
	}
	remainder = (sum * 10) % 11
	if remainder == 10 {
		remainder = 0
	}
	return remainder == int(cpf[10]-'0')
}

func validateCNPJ(cnpj string) bool {
	if matched, _ := regexp.MatchString(`^[0-9]{14}$`, cnpj); !matched {
		return false
	}

	weights1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weights2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	// Calcula o primeiro dígito verificador
	sum := 0
	for i := 0; i < 12; i++ {
		digit, _ := strconv.Atoi(string(cnpj[i]))
		sum += digit * weights1[i]
	}
	remainder := sum % 11
	if remainder < 2 {
		remainder = 0
	} else {
		remainder = 11 - remainder
	}
	if remainder != int(cnpj[12]-'0') {
		return false
	}

	// Calcula o segundo dígito verificador
	sum = 0
	for i := 0; i < 13; i++ {
		digit, _ := strconv.Atoi(string(cnpj[i]))
		sum += digit * weights2[i]
	}
	remainder = sum % 11
	if remainder < 2 {
		remainder = 0
	} else {
		remainder = 11 - remainder
	}
	return remainder == int(cnpj[13]-'0')
}
