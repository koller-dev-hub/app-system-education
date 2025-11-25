package school_event

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewSchoolCreatedEvent(t *testing.T) {
	schoolID := "123"
	name := "Test School"
	code := "TS001"
	address := "123 Test St"
	city := "Test City"
	state := "TS"
	zipCode := "12345"
	country := "Test Country"
	phoneNumber := "1234567890"
	email := "test@school.com"
	isActive := true
	description := "A test school"

	event := NewSchoolCreatedEvent(schoolID, name, code, address, city, state, zipCode, country, phoneNumber, email, isActive, description)

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
	assert.False(t, event.Date.IsZero())
	assert.WithinDuration(t, time.Now(), event.Date, time.Second)
}

func TestSchoolCreatedEvent_EventName(t *testing.T) {
	event := &SchoolCreatedEvent{}
	assert.Equal(t, "school.created", event.EventName())
}

func TestSchoolCreatedEvent_OccurredOn(t *testing.T) {
	now := time.Now()
	event := &SchoolCreatedEvent{Date: now}
	assert.Equal(t, now, event.OccurredOn())
}
