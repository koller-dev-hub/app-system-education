package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamkoller/system-education/shared/utils"
)

func TestIsValidCPF(t *testing.T) {
	tests := []struct {
		name     string
		cpf      string
		expected bool
	}{
		{"Valid CPF 1", "12345678909", true},     // Mathematically valid
		{"Valid CPF 2", "11144477735", true},     // Valid
		{"Valid CPF 3", "000.000.000-00", false}, // All same digits
		{"Valid CPF with formatting", "111.444.777-35", true},
		{"Invalid Length", "123", false},
		{"Invalid Check Digits", "11144477736", false},
		{"All same digits 1", "11111111111", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, utils.IsValidCPF(tt.cpf))
		})
	}
}

func TestFormatCPF(t *testing.T) {
	tests := []struct {
		name     string
		cpf      string
		expected string
	}{
		{"Clean CPF", "11144477735", "111.444.777-35"},
		{"Already Formatted", "111.444.777-35", "111.444.777-35"},
		{"Mixed Chars", "111a444b777c35", "111.444.777-35"},
		{"Invalid Length", "123", "123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, utils.FormatCPF(tt.cpf))
		})
	}
}

func TestCleanCPF(t *testing.T) {
	tests := []struct {
		name     string
		cpf      string
		expected string
	}{
		{"Only Digits", "12345678901", "12345678901"},
		{"With Formatting", "123.456.789-01", "12345678901"},
		{"With Letters", "123a456b", "123456"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, utils.CleanCPF(tt.cpf))
		})
	}
}
