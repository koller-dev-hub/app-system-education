package student_dtos

import "time"

type UpdateStudentDto struct {
	FullName       *string    `json:"full_name"`
	EnrollmentCode *string    `json:"enrollment_code"`
	Email          *string    `json:"email"`
	PhoneNumber    *string    `json:"phone_number"`
	DateOfBirth    *time.Time `json:"date_of_birth"`
	CPF            *string    `json:"cpf"`
	RG             *string    `json:"rg"`

	Address *string `json:"address"`
	City    *string `json:"city"`
	State   *string `json:"state"`
	ZipCode *string `json:"zip_code"`
	Country *string `json:"country"`

	SchoolID       *string    `json:"school_id"`
	SchoolName     *string    `json:"school_name"`
	SchoolCode     *string    `json:"school_code"`
	Grade          *string    `json:"grade"`
	ClassRoom      *string    `json:"class_room"`
	Shift          *string    `json:"shift"`
	EnrollmentDate *time.Time `json:"enrollment_date"`

	GuardianName  *string `json:"guardian_name"`
	GuardianPhone *string `json:"guardian_phone"`
	GuardianEmail *string `json:"guardian_email"`
	GuardianCPF   *string `json:"guardian_cpf"`

	IsActive     *bool   `json:"is_active"`
	Observations *string `json:"observations"`
}
