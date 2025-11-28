package student_mapper

import (
	"time"

	student_entity "github.com/williamkoller/system-education/internal/student/domain/entity"
)

type StudentResponse struct {
	ID             string    `json:"id"`
	FullName       string    `json:"full_name"`
	EnrollmentCode string    `json:"enrollment_code"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phone_number"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	CPF            string    `json:"cpf"`
	RG             string    `json:"rg"`
	Address        string    `json:"address"`
	City           string    `json:"city"`
	State          string    `json:"state"`
	ZipCode        string    `json:"zip_code"`
	Country        string    `json:"country"`
	SchoolID       string    `json:"school_id"`
	SchoolName     string    `json:"school_name"`
	SchoolCode     string    `json:"school_code"`
	Grade          string    `json:"grade"`
	ClassRoom      string    `json:"class_room"`
	Shift          string    `json:"shift"`
	EnrollmentDate time.Time `json:"enrollment_date"`
	GuardianName   string    `json:"guardian_name"`
	GuardianPhone  string    `json:"guardian_phone"`
	GuardianEmail  string    `json:"guardian_email"`
	GuardianCPF    string    `json:"guardian_cpf"`
	IsActive       bool      `json:"is_active"`
	Observations   string    `json:"observations"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ToStudentResponse(student *student_entity.Student) *StudentResponse {
	return &StudentResponse{
		ID:             student.ID,
		FullName:       student.PersonalInfo.FullName,
		EnrollmentCode: student.PersonalInfo.EnrollmentCode,
		Email:          student.PersonalInfo.Email,
		PhoneNumber:    student.PersonalInfo.PhoneNumber,
		DateOfBirth:    student.PersonalInfo.DateOfBirth,
		CPF:            student.PersonalInfo.CPF,
		RG:             student.PersonalInfo.RG,
		Address:        student.Address.Address,
		City:           student.Address.City,
		State:          student.Address.State,
		ZipCode:        student.Address.ZipCode,
		Country:        student.Address.Country,
		SchoolID:       student.School.SchoolID,
		SchoolName:     student.School.SchoolName,
		SchoolCode:     student.School.SchoolCode,
		Grade:          student.School.Grade,
		ClassRoom:      student.School.ClassRoom,
		Shift:          string(student.School.Shift),
		EnrollmentDate: student.School.EnrollmentDate,
		GuardianName:   student.Guardian.Name,
		GuardianPhone:  student.Guardian.Phone,
		GuardianEmail:  student.Guardian.Email,
		GuardianCPF:    student.Guardian.CPF,
		IsActive:       student.IsActive,
		Observations:   student.Observations,
		CreatedAt:      student.CreatedAt,
		UpdatedAt:      student.UpdatedAt,
	}
}

func ToStudentResponses(students []*student_entity.Student) []*StudentResponse {
	responses := make([]*StudentResponse, 0, len(students))
	for _, student := range students {
		responses = append(responses, ToStudentResponse(student))
	}
	return responses
}
