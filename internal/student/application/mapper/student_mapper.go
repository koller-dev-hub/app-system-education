package student_mapper

import (
	"time"

	student_entity "github.com/williamkoller/system-education/internal/student/domain/entity"
)

type StudentResponse struct {
	ID             string    `json:"id"`
	FullName       string    `json:"fullName"`
	EnrollmentCode string    `json:"enrollmentCode"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phoneNumber"`
	DateOfBirth    time.Time `json:"dateOfBirth"`
	CPF            string    `json:"cpf"`
	RG             string    `json:"rg"`
	Address        string    `json:"address"`
	City           string    `json:"city"`
	State          string    `json:"state"`
	ZipCode        string    `json:"zipCode"`
	Country        string    `json:"country"`
	SchoolID       string    `json:"schoolId"`
	SchoolName     string    `json:"schoolName"`
	SchoolCode     string    `json:"schoolCode"`
	Grade          string    `json:"grade"`
	ClassRoom      string    `json:"classRoom"`
	Shift          string    `json:"shift"`
	EnrollmentDate time.Time `json:"enrollmentDate"`
	GuardianName   string    `json:"guardianName"`
	GuardianPhone  string    `json:"guardianPhone"`
	GuardianEmail  string    `json:"guardianEmail"`
	GuardianCPF    string    `json:"guardianCpf"`
	IsActive       bool      `json:"isActive"`
	Observations   string    `json:"observations"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
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
