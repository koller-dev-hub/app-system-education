package student_entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/williamkoller/system-education/shared/utils"
)

type ValidationError struct {
	Errors []string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("validation failed: %s", strings.Join(v.Errors, ", "))
}

func ValidationStudent(s *Student) (*Student, error) {
	var errs []string

	// Personal Info Validation
	if strings.TrimSpace(s.PersonalInfo.FullName) == "" {
		errs = append(errs, "full name is required")
	}

	if !utils.IsValidCPF(s.PersonalInfo.CPF) {
		errs = append(errs, "invalid student cpf")
	} else {
		s.PersonalInfo.CPF = utils.FormatCPF(s.PersonalInfo.CPF)
	}

	if strings.TrimSpace(s.PersonalInfo.Email) == "" {
		errs = append(errs, "email is required")
	} else if !strings.Contains(s.PersonalInfo.Email, "@") {
		errs = append(errs, "invalid student email")
	}

	if s.PersonalInfo.DateOfBirth.After(time.Now()) {
		errs = append(errs, "date of birth cannot be in the future")
	}

	// Address Validation
	if strings.TrimSpace(s.Address.Address) == "" {
		errs = append(errs, "address is required")
	}
	if strings.TrimSpace(s.Address.City) == "" {
		errs = append(errs, "city is required")
	}
	if strings.TrimSpace(s.Address.State) == "" {
		errs = append(errs, "state is required")
	}
	if strings.TrimSpace(s.Address.ZipCode) == "" {
		errs = append(errs, "zip code is required")
	}
	if strings.TrimSpace(s.Address.Country) == "" {
		errs = append(errs, "country is required")
	}

	// School Info Validation
	if strings.TrimSpace(s.School.SchoolID) == "" {
		errs = append(errs, "school id is required")
	}

	switch s.School.Shift {
	case StudentShiftMorning, StudentShiftAfternoon, StudentShiftEvening:
		// valid
	default:
		errs = append(errs, "invalid shift")
	}

	// Guardian Info Validation
	if strings.TrimSpace(s.Guardian.Name) == "" {
		errs = append(errs, "guardian name is required")
	}

	if !utils.IsValidCPF(s.Guardian.CPF) {
		errs = append(errs, "invalid guardian cpf")
	} else {
		s.Guardian.CPF = utils.FormatCPF(s.Guardian.CPF)
	}

	if len(errs) > 0 {
		return nil, &ValidationError{Errors: errs}
	}

	return s, nil
}

func ValidationUpdateStudent(s *Student) (*Student, error) {
	var errs []string

	// Check format fields (Email, CPF) if they are present
	// Since this function receives a full Student object (updated), we check the fields on it.

	if !utils.IsValidCPF(s.PersonalInfo.CPF) {
		errs = append(errs, "invalid student cpf")
	} else {
		s.PersonalInfo.CPF = utils.FormatCPF(s.PersonalInfo.CPF)
	}

	if !strings.Contains(s.PersonalInfo.Email, "@") {
		errs = append(errs, "invalid student email")
	}

	if !utils.IsValidCPF(s.Guardian.CPF) {
		errs = append(errs, "invalid guardian cpf")
	} else {
		s.Guardian.CPF = utils.FormatCPF(s.Guardian.CPF)
	}

	if len(errs) > 0 {
		return nil, &ValidationError{Errors: errs}
	}

	// Delegate to full validation for required fields
	return ValidationStudent(s)
}
