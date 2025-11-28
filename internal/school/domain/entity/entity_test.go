package school_entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSchool(t *testing.T) {
	tests := []struct {
		name          string
		inputSchool   *School
		expectedError string
	}{
		{
			name: "Success",
			inputSchool: &School{
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
			name: "Validation Failure",
			inputSchool: &School{
				Name: "", // Missing name
			},
			expectedError: "validation failed: name is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			school, err := NewSchool(tt.inputSchool)

			if tt.expectedError == "" {
				assert.NoError(t, err)
				assert.NotNil(t, school)
				assert.Equal(t, tt.inputSchool.Name, school.Name)
				assert.Equal(t, tt.inputSchool.Code, school.Code)

				// Verify Domain Event
				events := school.PullDomainEvents()
				assert.Len(t, events, 1)
				assert.Equal(t, "school.created", events[0].EventName())
			} else {
				assert.Error(t, err)
				assert.Nil(t, school)
				assert.Contains(t, err.Error(), tt.expectedError)
			}
		})
	}
}

func TestPullDomainEvents(t *testing.T) {
	t.Run("Nil Receiver", func(t *testing.T) {
		var s *School
		events := s.PullDomainEvents()
		assert.Nil(t, events)
	})

	t.Run("Empty Events", func(t *testing.T) {
		s := &School{}
		events := s.PullDomainEvents()
		assert.Empty(t, events)
	})
}

func TestUpdateSchool(t *testing.T) {
	t.Run("should update school successfully", func(t *testing.T) {
		school := &School{
			ID:          "123",
			Name:        "Old Name",
			Code:        "OLD001",
			Address:     "Old Address",
			City:        "Old City",
			State:       "OC",
			ZipCode:     "11111",
			Country:     "Old Country",
			PhoneNumber: "1111111111",
			Email:       "old@school.com",
			IsActive:    true,
			Description: "Old description",
		}

		name := "New Name"
		code := "NEW001"
		address := "New Address"
		city := "New City"
		state := "NC"
		zipCode := "22222"
		country := "New Country"
		phoneNumber := "2222222222"
		email := "new@school.com"
		isActive := true
		description := "New description"

		err := school.UpdateSchool(&name, &code, &address, &city, &state, &zipCode, &country, &phoneNumber, &email, &isActive, &description)

		assert.NoError(t, err)
		assert.Equal(t, "New Name", school.Name)
		assert.Equal(t, "NEW001", school.Code)
		assert.Equal(t, "New Address", school.Address)
		assert.Equal(t, "New City", school.City)
		assert.Equal(t, "NC", school.State)
		assert.Equal(t, "22222", school.ZipCode)
		assert.Equal(t, "New Country", school.Country)
		assert.Equal(t, "2222222222", school.PhoneNumber)
		assert.Equal(t, "new@school.com", school.Email)
		assert.True(t, school.IsActive)
		assert.Equal(t, "New description", school.Description)
		assert.False(t, school.UpdatedAt.IsZero())

		// Verify Domain Event
		events := school.PullDomainEvents()
		assert.Len(t, events, 1)
		assert.Equal(t, "school.updated", events[0].EventName())
	})

	t.Run("should return error when validation fails - empty name", func(t *testing.T) {
		school := &School{
			ID:          "123",
			Name:        "Valid Name",
			Code:        "CODE",
			Address:     "Address",
			City:        "City",
			State:       "ST",
			ZipCode:     "12345",
			Country:     "Country",
			PhoneNumber: "1234567890",
			Email:       "test@school.com",
			IsActive:    true,
			Description: "Description",
		}

		name := ""
		code := "CODE"
		address := "Address"
		city := "City"
		state := "ST"
		zipCode := "12345"
		country := "Country"
		phoneNumber := "1234567890"
		email := "test@school.com"
		isActive := true
		description := "Description"

		err := school.UpdateSchool(&name, &code, &address, &city, &state, &zipCode, &country, &phoneNumber, &email, &isActive, &description)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "name is required")
	})

	t.Run("should return error when validation fails - invalid email", func(t *testing.T) {
		school := &School{
			ID:          "123",
			Name:        "Valid Name",
			Code:        "CODE",
			Address:     "Address",
			City:        "City",
			State:       "ST",
			ZipCode:     "12345",
			Country:     "Country",
			PhoneNumber: "1234567890",
			Email:       "test@school.com",
			IsActive:    true,
			Description: "Description",
		}

		name := "Valid Name"
		code := "CODE"
		address := "Address"
		city := "City"
		state := "ST"
		zipCode := "12345"
		country := "Country"
		phoneNumber := "1234567890"
		email := "invalidemail"
		isActive := true
		description := "Description"

		err := school.UpdateSchool(&name, &code, &address, &city, &state, &zipCode, &country, &phoneNumber, &email, &isActive, &description)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "email is invalid")
	})

	t.Run("should update school successfully - set isActive false", func(t *testing.T) {
		school := &School{
			ID:          "123",
			Name:        "Valid Name",
			Code:        "CODE",
			Address:     "Address",
			City:        "City",
			State:       "ST",
			ZipCode:     "12345",
			Country:     "Country",
			PhoneNumber: "1234567890",
			Email:       "test@school.com",
			IsActive:    true,
			Description: "Description",
		}

		name := "Valid Name"
		code := "CODE"
		address := "Address"
		city := "City"
		state := "ST"
		zipCode := "12345"
		country := "Country"
		phoneNumber := "1234567890"
		email := "test@school.com"
		isActive := false
		description := "Description"

		err := school.UpdateSchool(&name, &code, &address, &city, &state, &zipCode, &country, &phoneNumber, &email, &isActive, &description)

		assert.NoError(t, err)
		assert.False(t, school.IsActive)
	})

	t.Run("should return error when validation fails - empty description", func(t *testing.T) {
		school := &School{
			ID:          "123",
			Name:        "Valid Name",
			Code:        "CODE",
			Address:     "Address",
			City:        "City",
			State:       "ST",
			ZipCode:     "12345",
			Country:     "Country",
			PhoneNumber: "1234567890",
			Email:       "test@school.com",
			IsActive:    true,
			Description: "Description",
		}

		name := "Valid Name"
		code := "CODE"
		address := "Address"
		city := "City"
		state := "ST"
		zipCode := "12345"
		country := "Country"
		phoneNumber := "1234567890"
		email := "test@school.com"
		isActive := true
		description := ""

		err := school.UpdateSchool(&name, &code, &address, &city, &state, &zipCode, &country, &phoneNumber, &email, &isActive, &description)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "description is required")
	})
}
