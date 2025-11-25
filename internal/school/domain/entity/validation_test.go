package school_entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationSchool(t *testing.T) {
	tests := []struct {
		name          string
		school        *School
		expectedError string
	}{
		{
			name: "Success",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "test@school.com",
			},
			expectedError: "",
		},
		{
			name: "Missing Name",
			school: &School{
				Name:        "",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "test@school.com",
			},
			expectedError: "validation failed: name is required",
		},
		{
			name: "Missing Code",
			school: &School{
				Name:        "Test School",
				Code:        "",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "test@school.com",
			},
			expectedError: "validation failed: code is required",
		},
		{
			name: "Missing Address",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "test@school.com",
			},
			expectedError: "validation failed: address is required",
		},
		{
			name: "Missing City",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "test@school.com",
			},
			expectedError: "validation failed: city is required",
		},
		{
			name: "Missing State",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "test@school.com",
			},
			expectedError: "validation failed: state is required",
		},
		{
			name: "Missing ZipCode",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "test@school.com",
			},
			expectedError: "validation failed: zip code is required",
		},
		{
			name: "Missing Country",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "",
				PhoneNumber: "1234567890",
				Email:       "test@school.com",
			},
			expectedError: "validation failed: country is required",
		},
		{
			name: "Missing PhoneNumber",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "",
				Email:       "test@school.com",
			},
			expectedError: "validation failed: phone number is required",
		},
		{
			name: "Missing Email",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "",
			},
			expectedError: "validation failed: email is required",
		},
		{
			name: "Multiple Missing Fields",
			school: &School{
				Name:        "",
				Code:        "",
				Address:     "",
				City:        "",
				State:       "",
				ZipCode:     "",
				Country:     "",
				PhoneNumber: "",
				Email:       "",
			},
			expectedError: "validation failed: name is required, code is required, address is required, city is required, state is required, zip code is required, country is required, phone number is required, email is required",
		},
		{
			name: "Trim Space Check",
			school: &School{
				Name:        "   ",
				Code:        "   ",
				Address:     "   ",
				City:        "   ",
				State:       "   ",
				ZipCode:     "   ",
				Country:     "   ",
				PhoneNumber: "   ",
				Email:       "   ",
			},
			expectedError: "validation failed: name is required, code is required, address is required, city is required, state is required, zip code is required, country is required, phone number is required, email is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ValidationSchool(tt.school)
			if tt.expectedError == "" {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.school, result)
			} else {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Equal(t, tt.expectedError, err.Error())
			}
		})
	}
}

func TestValidationUpdateSchool(t *testing.T) {
	tests := []struct {
		name          string
		school        *School
		expectedError string
	}{
		{
			name: "Success",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "test@school.com",
			},
			expectedError: "",
		},
		{
			name: "Missing Name",
			school: &School{
				Name:        "",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "test@school.com",
			},
			expectedError: "validation failed: name is required",
		},
		{
			name: "Missing Code",
			school: &School{
				Name:        "Test School",
				Code:        "",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "test@school.com",
			},
			expectedError: "validation failed: code is required",
		},
		{
			name: "Missing Address",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "test@school.com",
			},
			expectedError: "validation failed: address is required",
		},
		{
			name: "Missing City",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "test@school.com",
			},
			expectedError: "validation failed: city is required",
		},
		{
			name: "Missing State",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "test@school.com",
			},
			expectedError: "validation failed: state is required",
		},
		{
			name: "Missing ZipCode",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "test@school.com",
			},
			expectedError: "validation failed: zip code is required",
		},
		{
			name: "Missing Country",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "",
				PhoneNumber: "1234567890",
				Email:       "test@school.com",
			},
			expectedError: "validation failed: country is required",
		},
		{
			name: "Missing PhoneNumber",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "",
				Email:       "test@school.com",
			},
			expectedError: "validation failed: phone number is required",
		},
		{
			name: "Missing Email",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "",
			},
			expectedError: "validation failed: email is required",
		},
		{
			name: "Multiple Missing Fields",
			school: &School{
				Name:        "",
				Code:        "",
				Address:     "",
				City:        "",
				State:       "",
				ZipCode:     "",
				Country:     "",
				PhoneNumber: "",
				Email:       "",
			},
			expectedError: "validation failed: name is required, code is required, address is required, city is required, state is required, zip code is required, country is required, phone number is required, email is required",
		},
		{
			name: "No Trim Space Check (Should Pass if spaces)",
			school: &School{
				Name:        "   ",
				Code:        "   ",
				Address:     "   ",
				City:        "   ",
				State:       "   ",
				ZipCode:     "   ",
				Country:     "   ",
				PhoneNumber: "   ",
				Email:       "   ",
			},
			expectedError: "", // ValidationUpdateSchool does NOT check TrimSpace in the provided code
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ValidationUpdateSchool(tt.school)
			if tt.expectedError == "" {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.school, result)
			} else {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Equal(t, tt.expectedError, err.Error())
			}
		})
	}
}
