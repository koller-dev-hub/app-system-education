package student_entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func createValidStudent() *Student {
	return &Student{
		PersonalInfo: PersonalInfo{
			FullName:       "John Doe",
			EnrollmentCode: "ST123",
			Email:          "john@example.com",
			PhoneNumber:    "123456789",
			DateOfBirth:    time.Now().AddDate(-10, 0, 0),
			CPF:            "11144477735",
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
	}
}

func TestNewStudent(t *testing.T) {
	s := createValidStudent()

	student, err := NewStudent(s)

	assert.NoError(t, err)
	assert.NotNil(t, student)
	assert.NotEmpty(t, student.ID)
	assert.Equal(t, "111.444.777-35", student.PersonalInfo.CPF)
}

func TestNewStudent_ValidationFailure(t *testing.T) {
	s := createValidStudent()
	s.PersonalInfo.FullName = "" // Invalid

	student, err := NewStudent(s)

	assert.Error(t, err)
	assert.Nil(t, student)
	assert.Contains(t, err.Error(), "full name is required")
}

func TestUpdate(t *testing.T) {
	s := createValidStudent()
	student, _ := NewStudent(s)
	originalUpdatedAt := student.UpdatedAt
	time.Sleep(1 * time.Millisecond) // Ensure time changes

	// Test Partial Update (Only Name) - Granular update check
	newFullName := "John Updated"

	err := student.Update(
		&newFullName,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)

	assert.NoError(t, err)
	assert.Equal(t, newFullName, student.PersonalInfo.FullName)
	assert.Equal(t, "john@example.com", student.PersonalInfo.Email) // Should remain unchanged!
	assert.Equal(t, student.Address.City, "New York")               // Address should not change
	assert.True(t, student.UpdatedAt.After(originalUpdatedAt))

	// Test Update Error (Invalid Data)
	invalidCPF := "123"

	err = student.Update(
		nil,
		nil,
		nil,
		nil,
		nil,
		&invalidCPF,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid student cpf")
}

func TestUpdate_FullCoverage(t *testing.T) {
	s := createValidStudent()
	student, _ := NewStudent(s)
	// originalUpdatedAt := student.UpdatedAt
	time.Sleep(1 * time.Millisecond)

	// Prepare full update data
	newFullName := "Updated Name"
	newEnrollmentCode := "ST999"
	newEmail := "updated@example.com"
	newPhoneNumber := "999999999"
	newDateOfBirth := time.Now().AddDate(-12, 0, 0)
	newCPF := "11144477735" // Valid but same for simplicity, logic checks != ""
	newRG := "RG999"

	newAddressStr := "Updated St"
	newCity := "Updated City"
	newState := "XX"
	newZipCode := "99999"
	newCountry := "Updated Country"

	newSchoolID := "SCH999"
	newSchoolName := "Updated School"
	newSchoolCode := "US"
	newGrade := "6"
	newClassRoom := "B"
	newShift := "afternoon"
	newEnrollmentDate := time.Now()

	newGuardianName := "Updated Guardian"
	newGuardianPhone := "888888888"
	newGuardianEmail := "guardian@updated.com"
	newGuardianCPF := "11144477735"

	isActive := false
	observations := "Updated Observations"

	err := student.Update(
		&newFullName,
		&newEnrollmentCode,
		&newEmail,
		&newPhoneNumber,
		&newDateOfBirth,
		&newCPF,
		&newRG,
		&newAddressStr,
		&newCity,
		&newState,
		&newZipCode,
		&newCountry,
		&newSchoolID,
		&newSchoolName,
		&newSchoolCode,
		&newGrade,
		&newClassRoom,
		&newShift,
		&newEnrollmentDate,
		&newGuardianName,
		&newGuardianPhone,
		&newGuardianEmail,
		&newGuardianCPF,
		&isActive,
		&observations,
	)

	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", student.PersonalInfo.FullName)
	assert.Equal(t, "ST999", student.PersonalInfo.EnrollmentCode)
	assert.Equal(t, "updated@example.com", student.PersonalInfo.Email)
	assert.Equal(t, "999999999", student.PersonalInfo.PhoneNumber)
	assert.Equal(t, "RG999", student.PersonalInfo.RG)

	assert.Equal(t, "Updated St", student.Address.Address)
	assert.Equal(t, "Updated City", student.Address.City)
	assert.Equal(t, "XX", student.Address.State)
	assert.Equal(t, "99999", student.Address.ZipCode)
	assert.Equal(t, "Updated Country", student.Address.Country)

	assert.Equal(t, "SCH999", student.School.SchoolID)
	assert.Equal(t, "Updated School", student.School.SchoolName)
	assert.Equal(t, "US", student.School.SchoolCode)
	assert.Equal(t, "6", student.School.Grade)
	assert.Equal(t, "B", student.School.ClassRoom)
	assert.Equal(t, StudentShiftAfternoon, student.School.Shift)

	assert.Equal(t, "Updated Guardian", student.Guardian.Name)
	assert.Equal(t, "888888888", student.Guardian.Phone)
	assert.Equal(t, "guardian@updated.com", student.Guardian.Email)

	assert.False(t, student.IsActive)
	assert.Equal(t, "Updated Observations", student.Observations)
}

func TestNewStudent_WithID(t *testing.T) {
	s := createValidStudent()
	s.ID = "existing-id"

	student, err := NewStudent(s)

	assert.NoError(t, err)
	assert.Equal(t, "existing-id", student.ID)
}
