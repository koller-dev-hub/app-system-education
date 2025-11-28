package student_entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidationStudent(t *testing.T) {
	validDateOfBirth := time.Now().AddDate(-10, 0, 0)
	validEnrollmentDate := time.Now()

	tests := []struct {
		name          string
		student       *Student
		expectedError string
	}{
		{
			name: "Success",
			student: &Student{
				PersonalInfo: PersonalInfo{
					FullName:    "John Doe",
					StudentID:   "ST123",
					Email:       "john@example.com",
					PhoneNumber: "123456789",
					DateOfBirth: validDateOfBirth,
					CPF:         "111.444.777-35",
				},
				Address: AddressInfo{
					Address: "123 Main St",
					City:    "New York",
					State:   "NY",
					ZipCode: "10001",
					Country: "USA",
				},
				School: SchoolInfo{
					SchoolID:       "SCH123",
					SchoolName:     "Test School",
					SchoolCode:     "TS",
					Grade:          "5",
					ClassRoom:      "A",
					Shift:          StudentShiftMorning,
					EnrollmentDate: validEnrollmentDate,
				},
				Guardian: GuardianInfo{
					Name:  "Jane Doe",
					Phone: "987654321",
					Email: "jane@example.com",
					CPF:   "11144477735",
				},
				IsActive: true,
			},
			expectedError: "",
		},
		{
			name: "Missing Required Fields",
			student: &Student{
				PersonalInfo: PersonalInfo{
					FullName:    "",
					CPF:         "invalid",
					Email:       "invalid",
					DateOfBirth: time.Now().Add(24 * time.Hour), // Future date
				},
				Address: AddressInfo{},
				School: SchoolInfo{
					Shift: "invalid",
				},
				Guardian: GuardianInfo{
					Name: "",
					CPF:  "invalid",
				},
			},
			expectedError: "validation failed: full name is required, invalid student cpf, invalid student email, date of birth cannot be in the future, address is required, city is required, state is required, zip code is required, country is required, school id is required, invalid shift, guardian name is required, invalid guardian cpf",
		},
		{
			name: "Invalid CPF",
			student: &Student{
				PersonalInfo: PersonalInfo{
					FullName:    "John Doe",
					Email:       "john@example.com",
					DateOfBirth: validDateOfBirth,
					CPF:         "12345678900",
				},
				Address: AddressInfo{
					Address: "123 Main St",
					City:    "New York",
					State:   "NY",
					ZipCode: "10001",
					Country: "USA",
				},
				School: SchoolInfo{
					SchoolID: "SCH123",
					Shift:    StudentShiftMorning,
				},
				Guardian: GuardianInfo{
					Name: "Jane Doe",
					CPF:  "11144477735",
				},
			},
			expectedError: "validation failed: invalid student cpf",
		},
		{
			name: "Invalid Guardian CPF in Update",
			student: &Student{
				PersonalInfo: PersonalInfo{
					FullName:    "John Doe",
					Email:       "john@example.com",
					DateOfBirth: time.Now().AddDate(-10, 0, 0),
					CPF:         "11144477735",
				},
				Address:  AddressInfo{Address: "Valid", City: "Valid", State: "Valid", ZipCode: "Valid", Country: "Valid"},
				School:   SchoolInfo{SchoolID: "Valid", Shift: StudentShiftMorning},
				Guardian: GuardianInfo{Name: "Valid", CPF: "123"},
			},
			expectedError: "validation failed: invalid guardian cpf",
		},
		{
			name: "Invalid Full Validation in Update (e.g. Missing Name)",
			student: &Student{
				PersonalInfo: PersonalInfo{
					FullName:    "", // Invalid (Missing)
					Email:       "john@example.com",
					DateOfBirth: time.Now().AddDate(-10, 0, 0),
					CPF:         "11144477735",
				},
				Address:  AddressInfo{Address: "Valid", City: "Valid", State: "Valid", ZipCode: "Valid", Country: "Valid"},
				School:   SchoolInfo{SchoolID: "Valid", Shift: StudentShiftMorning},
				Guardian: GuardianInfo{Name: "Valid", CPF: "11144477735"},
			},
			expectedError: "validation failed: full name is required",
		},
		{
			name: "Invalid Email Format (Missing @)",
			student: &Student{
				PersonalInfo: PersonalInfo{
					FullName:    "John Doe",
					Email:       "johnexample.com", // Invalid missing @
					DateOfBirth: validDateOfBirth,
					CPF:         "111.444.777-35",
				},
				Address:  AddressInfo{Address: "V", City: "V", State: "V", ZipCode: "V", Country: "V"},
				School:   SchoolInfo{SchoolID: "S", Shift: StudentShiftMorning},
				Guardian: GuardianInfo{Name: "G", CPF: "11144477735"},
			},
			expectedError: "validation failed: invalid student email",
		},
		{
			name: "Valid Shift",
			student: &Student{
				PersonalInfo: PersonalInfo{FullName: "J", Email: "j@e.com", DateOfBirth: validDateOfBirth, CPF: "11144477735"},
				Address:      AddressInfo{Address: "V", City: "V", State: "V", ZipCode: "V", Country: "V"},
				School:       SchoolInfo{SchoolID: "S", Shift: StudentShiftAfternoon}, // Valid
				Guardian:     GuardianInfo{Name: "G", CPF: "11144477735"},
			},
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ValidationStudent(tt.student)
			if tt.expectedError == "" {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				// Check CPF Formatting
				assert.Equal(t, "111.444.777-35", result.PersonalInfo.CPF)
				assert.Equal(t, "111.444.777-35", result.Guardian.CPF)
			} else {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Equal(t, tt.expectedError, err.Error())
			}
		})
	}
}

func TestValidationUpdateStudent(t *testing.T) {
	tests := []struct {
		name          string
		student       *Student
		expectedError string
	}{
		{
			name: "Success",
			student: &Student{
				PersonalInfo: PersonalInfo{
					FullName:    "John Doe",
					StudentID:   "ST123",
					Email:       "john@example.com",
					PhoneNumber: "123456789",
					DateOfBirth: time.Now().AddDate(-10, 0, 0),
					CPF:         "111.444.777-35",
				},
				Address: AddressInfo{
					Address: "123 Main St",
					City:    "New York",
					State:   "NY",
					ZipCode: "10001",
					Country: "USA",
				},
				School: SchoolInfo{
					SchoolID:       "SCH123",
					SchoolName:     "Test School",
					SchoolCode:     "TS",
					Grade:          "5",
					ClassRoom:      "A",
					Shift:          StudentShiftMorning,
					EnrollmentDate: time.Now(),
				},
				Guardian: GuardianInfo{
					Name:  "Jane Doe",
					Phone: "987654321",
					Email: "jane@example.com",
					CPF:   "11144477735",
				},
				IsActive: true,
			},
			expectedError: "",
		},
		{
			name: "Invalid Email in Update",
			student: &Student{
				PersonalInfo: PersonalInfo{
					FullName:    "John Doe",
					Email:       "invalid-email", // Invalid
					DateOfBirth: time.Now().AddDate(-10, 0, 0),
					CPF:         "111.444.777-35",
				},
				// Other fields valid...
				Address: AddressInfo{Address: "Valid", City: "Valid", State: "Valid", ZipCode: "Valid", Country: "Valid"},
				School:  SchoolInfo{SchoolID: "Valid", Shift: StudentShiftMorning},
				Guardian: GuardianInfo{
					Name: "Valid",
					CPF:  "11144477735",
				},
			},
			expectedError: "validation failed: invalid student email",
		},
		{
			name: "Invalid CPF in Update",
			student: &Student{
				PersonalInfo: PersonalInfo{
					FullName:    "John Doe",
					Email:       "john@example.com",
					DateOfBirth: time.Now().AddDate(-10, 0, 0),
					CPF:         "123", // Invalid
				},
				Address:  AddressInfo{Address: "Valid", City: "Valid", State: "Valid", ZipCode: "Valid", Country: "Valid"},
				School:   SchoolInfo{SchoolID: "Valid", Shift: StudentShiftMorning},
				Guardian: GuardianInfo{Name: "Valid", CPF: "11144477735"},
			},
			expectedError: "validation failed: invalid student cpf",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ValidationUpdateStudent(tt.student)
			if tt.expectedError == "" {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			} else {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Equal(t, tt.expectedError, err.Error())
			}
		})
	}
}
