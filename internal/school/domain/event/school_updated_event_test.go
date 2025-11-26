package school_event

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewSchoolUpdatedEvent(t *testing.T) {
	t.Run("should create school updated event successfully", func(t *testing.T) {
		schoolID := "123"
		name := "Updated School"
		code := "US001"
		address := "456 Updated St"
		city := "Updated City"
		state := "UC"
		zipCode := "54321"
		country := "Updated Country"
		phoneNumber := "0987654321"
		email := "updated@school.com"
		isActive := true
		description := "An updated school"

		event := NewSchoolUpdatedEvent(schoolID, name, code, address, city, state, zipCode, country, phoneNumber, email, isActive, description)

		assert.NotNil(t, event)
		assert.Equal(t, schoolID, event.SchoolID)
		assert.Equal(t, name, event.Name)
		assert.Equal(t, code, event.Code)
		assert.Equal(t, address, event.Address)
		assert.Equal(t, city, event.City)
		assert.Equal(t, state, event.State)
		assert.Equal(t, zipCode, event.ZipCode)
		assert.Equal(t, country, event.Country)
		assert.Equal(t, phoneNumber, event.PhoneNumber)
		assert.Equal(t, email, event.Email)
		assert.Equal(t, isActive, event.IsActive)
		assert.Equal(t, description, event.Description)
		assert.WithinDuration(t, time.Now(), event.Date, time.Second)
	})

	t.Run("should create event with empty fields", func(t *testing.T) {
		event := NewSchoolUpdatedEvent("", "", "", "", "", "", "", "", "", "", false, "")

		assert.NotNil(t, event)
		assert.Equal(t, "", event.SchoolID)
		assert.Equal(t, "", event.Name)
		assert.Equal(t, "", event.Code)
		assert.Equal(t, false, event.IsActive)
		assert.WithinDuration(t, time.Now(), event.Date, time.Second)
	})
}

func TestSchoolUpdatedEvent_EventName(t *testing.T) {
	t.Run("should return correct event name", func(t *testing.T) {
		event := NewSchoolUpdatedEvent("123", "School", "S001", "Address", "City", "ST", "12345", "Country", "1234567890", "test@school.com", true, "Description")

		eventName := event.EventName()

		assert.Equal(t, "school.updated", eventName)
	})
}

func TestSchoolUpdatedEvent_OccurredOn(t *testing.T) {
	t.Run("should return event date", func(t *testing.T) {
		event := NewSchoolUpdatedEvent("123", "School", "S001", "Address", "City", "ST", "12345", "Country", "1234567890", "test@school.com", true, "Description")

		occurredOn := event.OccurredOn()

		assert.Equal(t, event.Date, occurredOn)
		assert.WithinDuration(t, time.Now(), occurredOn, time.Second)
	})

	t.Run("should return consistent date", func(t *testing.T) {
		event := NewSchoolUpdatedEvent("123", "School", "S001", "Address", "City", "ST", "12345", "Country", "1234567890", "test@school.com", true, "Description")

		occurredOn1 := event.OccurredOn()
		time.Sleep(10 * time.Millisecond)
		occurredOn2 := event.OccurredOn()

		assert.Equal(t, occurredOn1, occurredOn2, "OccurredOn should return the same date on multiple calls")
	})
}

func TestSchoolUpdatedEvent_AllFields(t *testing.T) {
	t.Run("should preserve all field values", func(t *testing.T) {
		testCases := []struct {
			name        string
			schoolID    string
			schoolName  string
			code        string
			address     string
			city        string
			state       string
			zipCode     string
			country     string
			phoneNumber string
			email       string
			isActive    bool
			description string
		}{
			{
				name:        "Full data",
				schoolID:    "abc123",
				schoolName:  "Test School",
				code:        "TS001",
				address:     "123 Main St",
				city:        "Test City",
				state:       "TC",
				zipCode:     "12345",
				country:     "Test Country",
				phoneNumber: "1234567890",
				email:       "test@example.com",
				isActive:    true,
				description: "Test description",
			},
			{
				name:        "Minimal data",
				schoolID:    "xyz",
				schoolName:  "Min",
				code:        "M",
				address:     "A",
				city:        "C",
				state:       "S",
				zipCode:     "Z",
				country:     "Co",
				phoneNumber: "P",
				email:       "E",
				isActive:    false,
				description: "D",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				event := NewSchoolUpdatedEvent(
					tc.schoolID,
					tc.schoolName,
					tc.code,
					tc.address,
					tc.city,
					tc.state,
					tc.zipCode,
					tc.country,
					tc.phoneNumber,
					tc.email,
					tc.isActive,
					tc.description,
				)

				assert.Equal(t, tc.schoolID, event.SchoolID)
				assert.Equal(t, tc.schoolName, event.Name)
				assert.Equal(t, tc.code, event.Code)
				assert.Equal(t, tc.address, event.Address)
				assert.Equal(t, tc.city, event.City)
				assert.Equal(t, tc.state, event.State)
				assert.Equal(t, tc.zipCode, event.ZipCode)
				assert.Equal(t, tc.country, event.Country)
				assert.Equal(t, tc.phoneNumber, event.PhoneNumber)
				assert.Equal(t, tc.email, event.Email)
				assert.Equal(t, tc.isActive, event.IsActive)
				assert.Equal(t, tc.description, event.Description)
			})
		}
	})
}
