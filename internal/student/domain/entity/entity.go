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
	fullName *string,
	enrollmentCode *string,
	email *string,
	phoneNumber *string,
	dateOfBirth *time.Time,
	cpf *string,
	rg *string,
	address *string,
	city *string,
	state *string,
	zipCode *string,
	country *string,
	schoolID *string,
	schoolName *string,
	schoolCode *string,
	grade *string,
	classRoom *string,
	shift *string,
	enrollmentDate *time.Time,
	guardianName *string,
	guardianPhone *string,
	guardianEmail *string,
	guardianCPF *string,
	isActive *bool,
	observations *string,
) error {
	// Personal Info
	if fullName != nil {
		s.PersonalInfo.FullName = *fullName
	}
	if enrollmentCode != nil {
		s.PersonalInfo.EnrollmentCode = *enrollmentCode
	}
	if email != nil {
		s.PersonalInfo.Email = *email
	}
	if phoneNumber != nil {
		s.PersonalInfo.PhoneNumber = *phoneNumber
	}
	if dateOfBirth != nil {
		s.PersonalInfo.DateOfBirth = *dateOfBirth
	}
	if cpf != nil {
		s.PersonalInfo.CPF = *cpf
	}
	if rg != nil {
		s.PersonalInfo.RG = *rg
	}

	// Address Info
	if address != nil {
		s.Address.Address = *address
	}
	if city != nil {
		s.Address.City = *city
	}
	if state != nil {
		s.Address.State = *state
	}
	if zipCode != nil {
		s.Address.ZipCode = *zipCode
	}
	if country != nil {
		s.Address.Country = *country
	}

	// School Info
	if schoolID != nil {
		s.School.SchoolID = *schoolID
	}
	if schoolName != nil {
		s.School.SchoolName = *schoolName
	}
	if schoolCode != nil {
		s.School.SchoolCode = *schoolCode
	}
	if grade != nil {
		s.School.Grade = *grade
	}
	if classRoom != nil {
		s.School.ClassRoom = *classRoom
	}
	if shift != nil {
		s.School.Shift = Shift(*shift)
	}
	if enrollmentDate != nil {
		s.School.EnrollmentDate = *enrollmentDate
	}

	// Guardian Info
	if guardianName != nil {
		s.Guardian.Name = *guardianName
	}
	if guardianPhone != nil {
		s.Guardian.Phone = *guardianPhone
	}
	if guardianEmail != nil {
		s.Guardian.Email = *guardianEmail
	}
	if guardianCPF != nil {
		s.Guardian.CPF = *guardianCPF
	}

	// Metadata
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
