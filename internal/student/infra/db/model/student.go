package student_model

import (
	"time"

	school_model "github.com/williamkoller/system-education/internal/school/infra/db/model"
	student_entity "github.com/williamkoller/system-education/internal/student/domain/entity"
)

type Student struct {
	ID string `gorm:"primaryKey"`

	// Personal Info
	FullName       string
	EnrollmentCode string
	Email          string
	PhoneNumber    string
	DateOfBirth    time.Time
	CPF            string
	RG             string

	// Address Info
	AddressZipCode string
	AddressCity    string
	AddressState   string
	AddressStreet  string
	AddressCountry string

	// School Info
	SchoolID       string
	School         *school_model.School `gorm:"foreignKey:SchoolID"`
	SchoolGrade    string
	SchoolClass    string
	SchoolShift    string
	EnrollmentDate time.Time

	// Guardian Info
	GuardianName  string
	GuardianPhone string
	GuardianEmail string
	GuardianCPF   string

	// Status and Metadata
	IsActive     bool
	Observations string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Student) TableName() string {
	return "students"
}

func ToEntity(m *Student) *student_entity.Student {
	if m == nil {
		return nil
	}

	var schoolName, schoolCode string
	if m.School != nil {
		schoolName = m.School.Name
		schoolCode = m.School.Code
	}

	return &student_entity.Student{
		ID: m.ID,
		PersonalInfo: student_entity.PersonalInfo{
			FullName:       m.FullName,
			EnrollmentCode: m.EnrollmentCode,
			Email:          m.Email,
			PhoneNumber:    m.PhoneNumber,
			DateOfBirth:    m.DateOfBirth,
			CPF:            m.CPF,
			RG:             m.RG,
		},
		Address: student_entity.AddressInfo{
			Address: m.AddressStreet,
			City:    m.AddressCity,
			State:   m.AddressState,
			ZipCode: m.AddressZipCode,
			Country: m.AddressCountry,
		},
		School: student_entity.SchoolInfo{
			SchoolID:       m.SchoolID,
			SchoolName:     schoolName,
			SchoolCode:     schoolCode,
			Grade:          m.SchoolGrade,
			ClassRoom:      m.SchoolClass,
			Shift:          student_entity.Shift(m.SchoolShift),
			EnrollmentDate: m.EnrollmentDate,
		},
		Guardian: student_entity.GuardianInfo{
			Name:  m.GuardianName,
			Phone: m.GuardianPhone,
			Email: m.GuardianEmail,
			CPF:   m.GuardianCPF,
		},
		IsActive:     m.IsActive,
		Observations: m.Observations,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func ToEntities(ms []*Student) []*student_entity.Student {
	entities := make([]*student_entity.Student, 0, len(ms))
	for _, m := range ms {
		entities = append(entities, ToEntity(m))
	}
	return entities
}

func FromEntity(s *student_entity.Student) *Student {
	if s == nil {
		return nil
	}
	return &Student{
		ID:             s.ID,
		FullName:       s.PersonalInfo.FullName,
		EnrollmentCode: s.PersonalInfo.EnrollmentCode,
		Email:          s.PersonalInfo.Email,
		PhoneNumber:    s.PersonalInfo.PhoneNumber,
		DateOfBirth:    s.PersonalInfo.DateOfBirth,
		CPF:            s.PersonalInfo.CPF,
		RG:             s.PersonalInfo.RG,
		AddressStreet:  s.Address.Address,
		AddressCity:    s.Address.City,
		AddressState:   s.Address.State,
		AddressZipCode: s.Address.ZipCode,
		AddressCountry: s.Address.Country,
		SchoolID:       s.School.SchoolID,
		// SchoolName & SchoolCode are not stored in Student table anymore
		SchoolGrade:    s.School.Grade,
		SchoolClass:    s.School.ClassRoom,
		SchoolShift:    string(s.School.Shift),
		EnrollmentDate: s.School.EnrollmentDate,
		GuardianName:   s.Guardian.Name,
		GuardianPhone:  s.Guardian.Phone,
		GuardianEmail:  s.Guardian.Email,
		GuardianCPF:    s.Guardian.CPF,
		IsActive:       s.IsActive,
		Observations:   s.Observations,
		CreatedAt:      s.CreatedAt,
		UpdatedAt:      s.UpdatedAt,
	}
}

func FromEntities(es []*student_entity.Student) []*Student {
	models := make([]*Student, 0, len(es))
	for _, e := range es {
		models = append(models, FromEntity(e))
	}
	return models
}
