package student_dtos

import "time"

type AddStudentDto struct {
	FullName       string    `json:"full_name" binding:"required"`
	EnrollmentCode string    `json:"enrollment_code" binding:"required"`
	Email          string    `json:"email" binding:"required,email"`
	PhoneNumber    string    `json:"phone_number" binding:"required"`
	DateOfBirth    time.Time `json:"date_of_birth" binding:"required"`
	CPF            string    `json:"cpf" binding:"required"`
	RG             string    `json:"rg" binding:"required"`

	Address string `json:"address" binding:"required"`
	City    string `json:"city" binding:"required"`
	State   string `json:"state" binding:"required"`
	ZipCode string `json:"zip_code" binding:"required"`
	Country string `json:"country" binding:"required"`

	SchoolID       string    `json:"school_id" binding:"required"`
	SchoolName     string    `json:"school_name" binding:"required"`
	SchoolCode     string    `json:"school_code" binding:"required"`
	Grade          string    `json:"grade" binding:"required"`
	ClassRoom      string    `json:"class_room" binding:"required"`
	Shift          string    `json:"shift" binding:"required"`
	EnrollmentDate time.Time `json:"enrollment_date" binding:"required"`

	GuardianName  string `json:"guardian_name" binding:"required"`
	GuardianPhone string `json:"guardian_phone" binding:"required"`
	GuardianEmail string `json:"guardian_email" binding:"required"`
	GuardianCPF   string `json:"guardian_cpf" binding:"required"`

	IsActive     bool   `json:"is_active"`
	Observations string `json:"observations"`
}
