package school_mapper

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	school_entity "github.com/williamkoller/system-education/internal/school/domain/entity"
)

func TestToSchoolResponse(t *testing.T) {
	t.Run("should map school entity to response successfully", func(t *testing.T) {
		now := time.Now()
		school := &school_entity.School{
			ID:          "123",
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
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		response := ToSchoolResponse(school)

		assert.NotNil(t, response)
		assert.Equal(t, school.ID, response.ID)
		assert.Equal(t, school.Name, response.Name)
		assert.Equal(t, school.Code, response.Code)
		assert.Equal(t, school.Address, response.Address)
		assert.Equal(t, school.City, response.City)
		assert.Equal(t, school.State, response.State)
		assert.Equal(t, school.ZipCode, response.ZipCode)
		assert.Equal(t, school.Country, response.Country)
		assert.Equal(t, school.PhoneNumber, response.PhoneNumber)
		assert.Equal(t, school.Email, response.Email)
		assert.Equal(t, school.IsActive, response.IsActive)
		assert.Equal(t, school.Description, response.Description)
		assert.Equal(t, school.CreatedAt, response.CreatedAt)
		assert.Equal(t, school.UpdatedAt, response.UpdatedAt)
	})

	t.Run("should map school entity with empty fields", func(t *testing.T) {
		school := &school_entity.School{
			ID:        "456",
			Name:      "",
			Code:      "",
			IsActive:  false,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}

		response := ToSchoolResponse(school)

		assert.NotNil(t, response)
		assert.Equal(t, "456", response.ID)
		assert.Equal(t, "", response.Name)
		assert.Equal(t, "", response.Code)
		assert.Equal(t, false, response.IsActive)
		assert.Equal(t, time.Time{}, response.CreatedAt)
		assert.Equal(t, time.Time{}, response.UpdatedAt)
	})
}

func TestToSchoolResponses(t *testing.T) {
	t.Run("should map multiple schools to responses", func(t *testing.T) {
		now := time.Now()
		schools := []*school_entity.School{
			{
				ID:          "1",
				Name:        "School 1",
				Code:        "S001",
				Address:     "Address 1",
				City:        "City 1",
				State:       "ST1",
				ZipCode:     "11111",
				Country:     "Country 1",
				PhoneNumber: "1111111111",
				Email:       "school1@test.com",
				IsActive:    true,
				Description: "School 1 description",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			{
				ID:          "2",
				Name:        "School 2",
				Code:        "S002",
				Address:     "Address 2",
				City:        "City 2",
				State:       "ST2",
				ZipCode:     "22222",
				Country:     "Country 2",
				PhoneNumber: "2222222222",
				Email:       "school2@test.com",
				IsActive:    false,
				Description: "School 2 description",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		}

		responses := ToSchoolResponses(schools)

		assert.NotNil(t, responses)
		assert.Len(t, responses, 2)

		// Verify first school
		assert.Equal(t, "1", responses[0].ID)
		assert.Equal(t, "School 1", responses[0].Name)
		assert.Equal(t, "S001", responses[0].Code)
		assert.Equal(t, "school1@test.com", responses[0].Email)
		assert.True(t, responses[0].IsActive)

		// Verify second school
		assert.Equal(t, "2", responses[1].ID)
		assert.Equal(t, "School 2", responses[1].Name)
		assert.Equal(t, "S002", responses[1].Code)
		assert.Equal(t, "school2@test.com", responses[1].Email)
		assert.False(t, responses[1].IsActive)
	})

	t.Run("should return empty slice for empty input", func(t *testing.T) {
		schools := []*school_entity.School{}

		responses := ToSchoolResponses(schools)

		assert.NotNil(t, responses)
		assert.Len(t, responses, 0)
	})

	t.Run("should handle nil slice", func(t *testing.T) {
		var schools []*school_entity.School

		responses := ToSchoolResponses(schools)

		assert.NotNil(t, responses)
		assert.Len(t, responses, 0)
	})

	t.Run("should map single school in slice", func(t *testing.T) {
		now := time.Now()
		schools := []*school_entity.School{
			{
				ID:          "single",
				Name:        "Single School",
				Code:        "SS001",
				Address:     "Single Address",
				City:        "Single City",
				State:       "SC",
				ZipCode:     "99999",
				Country:     "Single Country",
				PhoneNumber: "9999999999",
				Email:       "single@test.com",
				IsActive:    true,
				Description: "Single school",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		}

		responses := ToSchoolResponses(schools)

		assert.NotNil(t, responses)
		assert.Len(t, responses, 1)
		assert.Equal(t, "single", responses[0].ID)
		assert.Equal(t, "Single School", responses[0].Name)
	})
}
