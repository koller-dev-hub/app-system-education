package student_usecase

import (
	"context"

	student_entity "github.com/williamkoller/system-education/internal/student/domain/entity"
	port_student_repository "github.com/williamkoller/system-education/internal/student/port/repository"
	port_student_usecase "github.com/williamkoller/system-education/internal/student/port/usecase"
	student_dtos "github.com/williamkoller/system-education/internal/student/presentation/dtos"
)

type StudentUsecase struct {
	repo port_student_repository.StudentRepository
}

func NewStudentUsecase(repo port_student_repository.StudentRepository) *StudentUsecase {
	return &StudentUsecase{repo: repo}
}

var _ port_student_usecase.StudentUsecase = &StudentUsecase{}

func (s *StudentUsecase) Create(ctx context.Context, input student_dtos.AddStudentDto) (*student_entity.Student, error) {
	student := &student_entity.Student{
		PersonalInfo: student_entity.PersonalInfo{
			FullName:       input.FullName,
			EnrollmentCode: input.EnrollmentCode,
			Email:          input.Email,
			PhoneNumber:    input.PhoneNumber,
			DateOfBirth:    input.DateOfBirth,
			CPF:            input.CPF,
			RG:             input.RG,
		},
		Address: student_entity.AddressInfo{
			Address: input.Address,
			City:    input.City,
			State:   input.State,
			ZipCode: input.ZipCode,
			Country: input.Country,
		},
		School: student_entity.SchoolInfo{
			SchoolID:       input.SchoolID,
			SchoolName:     input.SchoolName,
			SchoolCode:     input.SchoolCode,
			Grade:          input.Grade,
			ClassRoom:      input.ClassRoom,
			Shift:          student_entity.Shift(input.Shift),
			EnrollmentDate: input.EnrollmentDate,
		},
		Guardian: student_entity.GuardianInfo{
			Name:  input.GuardianName,
			Phone: input.GuardianPhone,
			Email: input.GuardianEmail,
			CPF:   input.GuardianCPF,
		},
		IsActive:     input.IsActive,
		Observations: input.Observations,
	}

	newStudent, err := student_entity.NewStudent(student)
	if err != nil {
		return nil, err
	}

	return s.repo.Save(ctx, newStudent)
}

func (s *StudentUsecase) FindAll(ctx context.Context) ([]*student_entity.Student, error) {
	return s.repo.FindAll(ctx)
}

func (s *StudentUsecase) FindById(ctx context.Context, id string) (*student_entity.Student, error) {
	return s.repo.FindById(ctx, id)
}

func (s *StudentUsecase) Update(ctx context.Context, id string, input student_dtos.UpdateStudentDto) (*student_entity.Student, error) {
	studentFound, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	err = studentFound.Update(
		input.FullName,
		input.EnrollmentCode,
		input.Email,
		input.PhoneNumber,
		input.DateOfBirth,
		input.CPF,
		input.RG,
		input.Address,
		input.City,
		input.State,
		input.ZipCode,
		input.Country,
		input.SchoolID,
		input.SchoolName,
		input.SchoolCode,
		input.Grade,
		input.ClassRoom,
		input.Shift,
		input.EnrollmentDate,
		input.GuardianName,
		input.GuardianPhone,
		input.GuardianEmail,
		input.GuardianCPF,
		input.IsActive,
		input.Observations,
	)

	if err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, id, studentFound)
}

func (s *StudentUsecase) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
