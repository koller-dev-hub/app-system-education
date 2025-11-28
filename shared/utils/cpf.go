package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// CPFRegex is the regular expression to validate CPF format
var CPFRegex = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)

// IsValidCPF checks if the provided CPF is valid
func IsValidCPF(cpf string) bool {
	cpf = CleanCPF(cpf)

	if len(cpf) != 11 {
		return false
	}

	if isAllSameDigits(cpf) {
		return false
	}

	// Validate first digit
	sum := 0
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (10 - i)
	}
	remainder := sum % 11
	digit1 := 0
	if remainder >= 2 {
		digit1 = 11 - remainder
	}

	if digit1 != int(cpf[9]-'0') {
		return false
	}

	// Validate second digit
	sum = 0
	for i := 0; i < 10; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (11 - i)
	}
	remainder = sum % 11
	digit2 := 0
	if remainder >= 2 {
		digit2 = 11 - remainder
	}

	return digit2 == int(cpf[10]-'0')
}

// FormatCPF formats the CPF to the standard format XXX.XXX.XXX-XX
func FormatCPF(cpf string) string {
	cpf = CleanCPF(cpf)
	if len(cpf) != 11 {
		return cpf
	}
	return fmt.Sprintf("%s.%s.%s-%s", cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:11])
}

// CleanCPF removes all non-digit characters from the string
func CleanCPF(cpf string) string {
	return strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, cpf)
}

func isAllSameDigits(cpf string) bool {
	first := cpf[0]
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != first {
			return false
		}
	}
	return true
}
