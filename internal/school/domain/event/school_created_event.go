package school_event

import "time"

type SchoolCreatedEvent struct {
	SchoolID    string
	Name        string
	Code        string
	Address     string
	City        string
	State       string
	ZipCode     string
	Country     string
	PhoneNumber string
	Email       string
	IsActive    bool
	Description string
	Date        time.Time
}

func NewSchoolCreatedEvent(schoolID string, name string, code string, address string, city string, state string, zipCode string, country string, phoneNumber string, email string, isActive bool, description string) *SchoolCreatedEvent {
	return &SchoolCreatedEvent{
		SchoolID:    schoolID,
		Name:        name,
		Code:        code,
		Address:     address,
		City:        city,
		State:       state,
		ZipCode:     zipCode,
		Country:     country,
		PhoneNumber: phoneNumber,
		Email:       email,
		IsActive:    isActive,
		Description: description,
		Date:        time.Now(),
	}
}

func (e *SchoolCreatedEvent) EventName() string {
	return "school.created"
}

func (e *SchoolCreatedEvent) OccurredOn() time.Time {
	return e.Date
}
