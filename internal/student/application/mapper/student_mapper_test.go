package student_mapper

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	student_entity "github.com/williamkoller/system-education/internal/student/domain/entity"
)

func TestToStudentResponse(t *testing.T) {
	t.Run("should map student entity to response successfully", func(t *testing.T) {
		now := time.Now()
		student := &student_entity.Student{
			ID: "123",
			PersonalInfo: student_entity.PersonalInfo{
				FullName:       "John Doe",
				EnrollmentCode: "2023001",
				Email:          "john@example.com",
				PhoneNumber:    "1234567890",
				DateOfBirth:    now.AddDate(-10, 0, 0),
				CPF:            "97093236014",
				RG:             "1234567",
			},
			Address: student_entity.AddressInfo{
				Address: "123 Main St",
				City:    "City",
				State:   "ST",
				ZipCode: "12345",
				Country: "Country",
			},
			School: student_entity.SchoolInfo{
				SchoolID:       "school-1",
				SchoolName:     "School Name",
				SchoolCode:     "SC001",
				Grade:          "5th",
				ClassRoom:      "A",
				Shift:          student_entity.StudentShiftMorning,
				EnrollmentDate: now,
			},
			Guardian: student_entity.GuardianInfo{
				Name:  "Guardian",
				Phone: "0987654321",
				Email: "guardian@example.com",
				CPF:   "97093236014",
			},
			IsActive:     true,
			Observations: "None",
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		response := ToStudentResponse(student)

		assert.NotNil(t, response)
		assert.Equal(t, student.ID, response.ID)

		// Verify JSON keys
		jsonData, err := json.Marshal(response)
		assert.NoError(t, err)

		jsonMap := make(map[string]interface{})
		err = json.Unmarshal(jsonData, &jsonMap)
		assert.NoError(t, err)

		assert.Contains(t, jsonMap, "fullName")
		assert.Contains(t, jsonMap, "enrollmentCode")
		assert.Contains(t, jsonMap, "phoneNumber")
		assert.Contains(t, jsonMap, "dateOfBirth")
		assert.Contains(t, jsonMap, "zipCode")
		assert.Contains(t, jsonMap, "schoolId")
		assert.Contains(t, jsonMap, "schoolName")
		assert.Contains(t, jsonMap, "schoolCode")
		assert.Contains(t, jsonMap, "classRoom")
		assert.Contains(t, jsonMap, "enrollmentDate")
		assert.Contains(t, jsonMap, "guardianName")
		assert.Contains(t, jsonMap, "guardianPhone")
		assert.Contains(t, jsonMap, "guardianEmail")
		assert.Contains(t, jsonMap, "guardianCpf")
		assert.Contains(t, jsonMap, "isActive")
		assert.Contains(t, jsonMap, "createdAt")
		assert.Contains(t, jsonMap, "updatedAt")

		assert.Equal(t, student.PersonalInfo.FullName, response.FullName)
		assert.Equal(t, student.PersonalInfo.EnrollmentCode, response.EnrollmentCode)
		assert.Equal(t, student.PersonalInfo.Email, response.Email)
		assert.Equal(t, student.PersonalInfo.PhoneNumber, response.PhoneNumber)
		assert.Equal(t, student.PersonalInfo.DateOfBirth, response.DateOfBirth)
		assert.Equal(t, student.PersonalInfo.CPF, response.CPF)
		assert.Equal(t, student.PersonalInfo.RG, response.RG)
		assert.Equal(t, student.Address.Address, response.Address)
		assert.Equal(t, student.Address.City, response.City)
		assert.Equal(t, student.Address.State, response.State)
		assert.Equal(t, student.Address.ZipCode, response.ZipCode)
		assert.Equal(t, student.Address.Country, response.Country)
		assert.Equal(t, student.School.SchoolID, response.SchoolID)
		assert.Equal(t, student.School.SchoolName, response.SchoolName)
		assert.Equal(t, student.School.SchoolCode, response.SchoolCode)
		assert.Equal(t, student.School.Grade, response.Grade)
		assert.Equal(t, student.School.ClassRoom, response.ClassRoom)
		assert.Equal(t, string(student.School.Shift), response.Shift)
		assert.Equal(t, student.School.EnrollmentDate, response.EnrollmentDate)
		assert.Equal(t, student.Guardian.Name, response.GuardianName)
		assert.Equal(t, student.Guardian.Phone, response.GuardianPhone)
		assert.Equal(t, student.Guardian.Email, response.GuardianEmail)
		assert.Equal(t, student.Guardian.CPF, response.GuardianCPF)
		assert.Equal(t, student.IsActive, response.IsActive)
		assert.Equal(t, student.Observations, response.Observations)
		assert.Equal(t, student.CreatedAt, response.CreatedAt)
		assert.Equal(t, student.UpdatedAt, response.UpdatedAt)
	})

	t.Run("should map student entity with empty fields", func(t *testing.T) {
		student := &student_entity.Student{
			ID: "456",
			PersonalInfo: student_entity.PersonalInfo{
				FullName: "",
			},
			IsActive: false,
		}

		response := ToStudentResponse(student)

		assert.NotNil(t, response)
		assert.Equal(t, "456", response.ID)
		assert.Equal(t, "", response.FullName)
		assert.Equal(t, false, response.IsActive)
	})
}

func TestToStudentResponses(t *testing.T) {
	t.Run("should map multiple students to responses", func(t *testing.T) {
		students := []*student_entity.Student{
			{
				ID: "1",
				PersonalInfo: student_entity.PersonalInfo{
					FullName: "Student 1",
				},
			},
			{
				ID: "2",
				PersonalInfo: student_entity.PersonalInfo{
					FullName: "Student 2",
				},
			},
		}

		responses := ToStudentResponses(students)

		assert.Len(t, responses, 2)
		assert.Equal(t, "1", responses[0].ID)
		assert.Equal(t, "Student 1", responses[0].FullName)
		assert.Equal(t, "2", responses[1].ID)
		assert.Equal(t, "Student 2", responses[1].FullName)
	})

	t.Run("should return empty slice when input is empty", func(t *testing.T) {
		var students []*student_entity.Student
		responses := ToStudentResponses(students)
		assert.NotNil(t, responses)
		assert.Empty(t, responses)
	})
}
