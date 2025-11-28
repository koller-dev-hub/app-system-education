package student_entity

import (
	"time"

	"github.com/google/uuid"
	shared_event "github.com/williamkoller/system-education/shared/domain/event"
)

type PersonalInfo struct {
	FullName       string
	EnrollmentCode string // Enrollment Code is unique for each student
	Email          string
	PhoneNumber    string
	DateOfBirth    time.Time
	CPF            string
	RG             string
}

type AddressInfo struct {
	Address string
	City    string
	State   string
	ZipCode string
	Country string
}

type Shift string

var (
	StudentShiftMorning   Shift = "morning"
	StudentShiftAfternoon Shift = "afternoon"
	StudentShiftEvening   Shift = "evening"
)

type SchoolInfo struct {
	SchoolID       string
	SchoolName     string
	SchoolCode     string
	Grade          string
	ClassRoom      string
	Shift          Shift
	EnrollmentDate time.Time
}

type GuardianInfo struct {
	Name  string
	Phone string
	Email string
	CPF   string
}

type Student struct {
	ID           string
	PersonalInfo PersonalInfo
	Address      AddressInfo
	School       SchoolInfo
	Guardian     GuardianInfo

	// Status and Metadata
	IsActive     bool
	Observations string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	shared_event.AggregateRoot
}

func NewStudent(s *Student) (*Student, error) {
	vs, err := ValidationStudent(s)
	if err != nil {
		return nil, err
	}

	id := s.ID
	if id == "" {
		id = uuid.New().String()
	}

	student := &Student{
		ID:           id,
		PersonalInfo: vs.PersonalInfo,
		Address:      vs.Address,
		School:       vs.School,
		Guardian:     vs.Guardian,
		IsActive:     vs.IsActive,
		Observations: vs.Observations,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	return student, nil
}

func (s *Student) Update(
	personalInfo *PersonalInfo,
	address *AddressInfo,
	school *SchoolInfo,
	guardian *GuardianInfo,
	isActive *bool,
	observations *string,
) error {
	if personalInfo != nil {
		if personalInfo.FullName != "" {
			s.PersonalInfo.FullName = personalInfo.FullName
		}
		if personalInfo.EnrollmentCode != "" {
			s.PersonalInfo.EnrollmentCode = personalInfo.EnrollmentCode
		}
		if personalInfo.Email != "" {
			s.PersonalInfo.Email = personalInfo.Email
		}
		if personalInfo.PhoneNumber != "" {
			s.PersonalInfo.PhoneNumber = personalInfo.PhoneNumber
		}
		if !personalInfo.DateOfBirth.IsZero() {
			s.PersonalInfo.DateOfBirth = personalInfo.DateOfBirth
		}
		if personalInfo.CPF != "" {
			s.PersonalInfo.CPF = personalInfo.CPF
		}
		if personalInfo.RG != "" {
			s.PersonalInfo.RG = personalInfo.RG
		}
	}

	if address != nil {
		if address.Address != "" {
			s.Address.Address = address.Address
		}
		if address.City != "" {
			s.Address.City = address.City
		}
		if address.State != "" {
			s.Address.State = address.State
		}
		if address.ZipCode != "" {
			s.Address.ZipCode = address.ZipCode
		}
		if address.Country != "" {
			s.Address.Country = address.Country
		}
	}

	if school != nil {
		if school.SchoolID != "" {
			s.School.SchoolID = school.SchoolID
		}
		if school.SchoolName != "" {
			s.School.SchoolName = school.SchoolName
		}
		if school.SchoolCode != "" {
			s.School.SchoolCode = school.SchoolCode
		}
		if school.Grade != "" {
			s.School.Grade = school.Grade
		}
		if school.ClassRoom != "" {
			s.School.ClassRoom = school.ClassRoom
		}
		if school.Shift != "" {
			s.School.Shift = school.Shift
		}
		if !school.EnrollmentDate.IsZero() {
			s.School.EnrollmentDate = school.EnrollmentDate
		}
	}

	if guardian != nil {
		if guardian.Name != "" {
			s.Guardian.Name = guardian.Name
		}
		if guardian.Phone != "" {
			s.Guardian.Phone = guardian.Phone
		}
		if guardian.Email != "" {
			s.Guardian.Email = guardian.Email
		}
		if guardian.CPF != "" {
			s.Guardian.CPF = guardian.CPF
		}
	}

	if isActive != nil {
		s.IsActive = *isActive
	}
	if observations != nil {
		s.Observations = *observations
	}
	s.UpdatedAt = time.Now()

	if _, err := ValidationUpdateStudent(s); err != nil {
		return err
	}

	return nil
}
