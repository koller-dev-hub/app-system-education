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
				IsActive:    true,
				Description: "A test school",
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
				IsActive:    true,
				Description: "A test school",
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
				IsActive:    true,
				Description: "A test school",
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
				IsActive:    true,
				Description: "A test school",
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
				IsActive:    true,
				Description: "A test school",
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
				IsActive:    true,
				Description: "A test school",
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
				IsActive:    true,
				Description: "A test school",
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
				IsActive:    true,
				Description: "A test school",
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
				IsActive:    true,
				Description: "A test school",
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
				IsActive:    true,
				Description: "A test school",
			},
			expectedError: "validation failed: email is required, email is invalid",
		},
		{
			name: "Invalid Email Format",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "invalidemail",
				IsActive:    true,
				Description: "A test school",
			},
			expectedError: "validation failed: email is invalid",
		},
		{
			name: "IsActive False",
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
				IsActive:    false,
				Description: "A test school",
			},
			expectedError: "validation failed: is active is required",
		},
		{
			name: "Missing Description",
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
				IsActive:    true,
				Description: "",
			},
			expectedError: "validation failed: description is required",
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
				IsActive:    false,
				Description: "",
			},
			expectedError: "validation failed: name is required, code is required, address is required, city is required, state is required, zip code is required, country is required, phone number is required, email is required, email is invalid, is active is required, description is required",
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
				IsActive:    false,
				Description: "   ",
			},
			expectedError: "validation failed: name is required, code is required, address is required, city is required, state is required, zip code is required, country is required, phone number is required, email is required, email is invalid, is active is required, description is required",
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
				IsActive:    true,
				Description: "A test school",
			},
			expectedError: "",
		},
		{
			name: "Invalid Email Format",
			school: &School{
				Name:        "Test School",
				Code:        "TS001",
				Address:     "123 Test St",
				City:        "Test City",
				State:       "TS",
				ZipCode:     "12345",
				Country:     "Test Country",
				PhoneNumber: "1234567890",
				Email:       "invalidemail",
				IsActive:    true,
				Description: "A test school",
			},
			expectedError: "validation failed: email is invalid",
		},
		{
			name: "Empty Name",
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
				IsActive:    true,
				Description: "A test school",
			},
			expectedError: "validation failed: name is required",
		},
		{
			name: "Empty Code",
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
				IsActive:    true,
				Description: "A test school",
			},
			expectedError: "validation failed: code is required",
		},
		{
			name: "Empty Address",
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
				IsActive:    true,
				Description: "A test school",
			},
			expectedError: "validation failed: address is required",
		},
		{
			name: "Empty City",
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
				IsActive:    true,
				Description: "A test school",
			},
			expectedError: "validation failed: city is required",
		},
		{
			name: "Empty State",
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
				IsActive:    true,
				Description: "A test school",
			},
			expectedError: "validation failed: state is required",
		},
		{
			name: "Empty ZipCode",
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
				IsActive:    true,
				Description: "A test school",
			},
			expectedError: "validation failed: zip code is required",
		},
		{
			name: "Empty Country",
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
				IsActive:    true,
				Description: "A test school",
			},
			expectedError: "validation failed: country is required",
		},
		{
			name: "Empty PhoneNumber",
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
				IsActive:    true,
				Description: "A test school",
			},
			expectedError: "validation failed: phone number is required",
		},
		{
			name: "Empty Email",
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
				IsActive:    true,
				Description: "A test school",
			},
			expectedError: "validation failed: email is required, email is invalid",
		},
		{
			name: "Empty Description",
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
				IsActive:    true,
				Description: "",
			},
			expectedError: "validation failed: description is required",
		},
		{
			name: "Multiple Empty Fields",
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
				IsActive:    true,
				Description: "",
			},
			expectedError: "validation failed: name is required, code is required, address is required, city is required, state is required, zip code is required, country is required, phone number is required, email is required, email is invalid, description is required",
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
