package school_entity

import (
	"time"

	school_event "github.com/williamkoller/system-education/internal/school/domain/event"
	shared_event "github.com/williamkoller/system-education/shared/domain/event"
)

type School struct {
	ID          string
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
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	shared_event.AggregateRoot
}

func NewSchool(s *School) (*School, error) {
	vs, err := ValidationSchool(s)
	if err != nil {
		return nil, err
	}

	school := &School{
		ID:          vs.ID,
		Name:        vs.Name,
		Code:        vs.Code,
		Address:     vs.Address,
		City:        vs.City,
		State:       vs.State,
		ZipCode:     vs.ZipCode,
		Country:     vs.Country,
		PhoneNumber: vs.PhoneNumber,
		Email:       vs.Email,
		IsActive:    vs.IsActive,
		Description: vs.Description,
		CreatedAt:   vs.CreatedAt,
		UpdatedAt:   vs.UpdatedAt,
		DeletedAt:   vs.DeletedAt,
	}

	school.AddDomainEvent(school_event.NewSchoolCreatedEvent(school.ID, school.Name, school.Code, school.Address, school.City, school.State, school.ZipCode, school.Country, school.PhoneNumber, school.Email, school.IsActive, school.Description))

	return school, nil
}

func (s *School) PullDomainEvents() []shared_event.Event {
	if s == nil {
		return nil
	}
	return s.AggregateRoot.PullDomainEvents()
}

func (s *School) UpdateSchool(name *string, code *string, address *string, city *string, state *string, zipCode *string, country *string, phoneNumber *string, email *string, isActive *bool, description *string) error {
	if name != nil {
		s.Name = *name
	}
	if code != nil {
		s.Code = *code
	}
	if address != nil {
		s.Address = *address
	}
	if city != nil {
		s.City = *city
	}
	if state != nil {
		s.State = *state
	}
	if zipCode != nil {
		s.ZipCode = *zipCode
	}
	if country != nil {
		s.Country = *country
	}
	if phoneNumber != nil {
		s.PhoneNumber = *phoneNumber
	}
	if email != nil {
		s.Email = *email
	}
	if isActive != nil {
		s.IsActive = *isActive
	}
	if description != nil {
		s.Description = *description
	}

	s.UpdatedAt = time.Now()

	vs, err := ValidationUpdateSchool(s)
	if err != nil {
		return err
	}

	s.AddDomainEvent(school_event.NewSchoolUpdatedEvent(s.ID, vs.Name, vs.Code, vs.Address, vs.City, vs.State, vs.ZipCode, vs.Country, vs.PhoneNumber, vs.Email, vs.IsActive, vs.Description))

	return nil
}
