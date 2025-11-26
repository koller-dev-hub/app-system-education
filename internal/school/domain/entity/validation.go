package school_entity

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Errors []string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("validation failed: %s", strings.Join(v.Errors, ", "))
}

func ValidationSchool(s *School) (*School, error) {
	var errs []string

	if strings.TrimSpace(s.Name) == "" {
		errs = append(errs, "name is required")
	}

	if strings.TrimSpace(s.Code) == "" {
		errs = append(errs, "code is required")
	}

	if strings.TrimSpace(s.Address) == "" {
		errs = append(errs, "address is required")
	}

	if strings.TrimSpace(s.City) == "" {
		errs = append(errs, "city is required")
	}

	if strings.TrimSpace(s.State) == "" {
		errs = append(errs, "state is required")
	}

	if strings.TrimSpace(s.ZipCode) == "" {
		errs = append(errs, "zip code is required")
	}

	if strings.TrimSpace(s.Country) == "" {
		errs = append(errs, "country is required")
	}

	if strings.TrimSpace(s.PhoneNumber) == "" {
		errs = append(errs, "phone number is required")
	}

	if strings.TrimSpace(s.Email) == "" {
		errs = append(errs, "email is required")
	}

	if !strings.Contains(s.Email, "@") {
		errs = append(errs, "email is invalid")
	}

	if !s.IsActive {
		errs = append(errs, "is active is required")
	}

	if strings.TrimSpace(s.Description) == "" {
		errs = append(errs, "description is required")
	}

	if len(errs) > 0 {
		return nil, &ValidationError{Errors: errs}
	}

	return s, nil
}

func ValidationUpdateSchool(s *School) (*School, error) {
	var errs []string

	if !strings.Contains(s.Email, "@") {
		errs = append(errs, "email is invalid")
	}

	if len(errs) > 0 {
		return nil, &ValidationError{Errors: errs}
	}

	return s, nil
}
