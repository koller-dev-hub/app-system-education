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
