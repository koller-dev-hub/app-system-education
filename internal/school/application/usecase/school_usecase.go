package school_usecase

import (
	"github.com/google/uuid"
	school_entity "github.com/williamkoller/system-education/internal/school/domain/entity"
	port_school_repository "github.com/williamkoller/system-education/internal/school/port/repository"
	port_school_usecase "github.com/williamkoller/system-education/internal/school/port/usecase"
	school_dtos "github.com/williamkoller/system-education/internal/school/presentation/dtos"
)

type SchoolUseCase struct {
	repo port_school_repository.SchoolRepository
}

func NewSchoolUseCase(repo port_school_repository.SchoolRepository) *SchoolUseCase {
	return &SchoolUseCase{
		repo: repo,
	}
}

var _ port_school_usecase.SchoolUseCase = &SchoolUseCase{}

func (s *SchoolUseCase) Create(input school_dtos.AddSchoolDto) (*school_entity.School, error) {
	school, err := school_entity.NewSchool(&school_entity.School{
		ID:          uuid.New().String(),
		Name:        input.Name,
		Code:        input.Code,
		Address:     input.Address,
		City:        input.City,
		State:       input.State,
		ZipCode:     input.ZipCode,
		Country:     input.Country,
		PhoneNumber: input.PhoneNumber,
		Email:       input.Email,
		IsActive:    input.IsActive,
		Description: input.Description,
	})

	if err != nil {
		return nil, err
	}

	school, err = s.repo.Save(school)
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (s *SchoolUseCase) FindAll() ([]*school_entity.School, error) {
	schools, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return schools, nil
}

func (s *SchoolUseCase) FindById(id string) (*school_entity.School, error) {
	school, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (s *SchoolUseCase) Update(id string, update school_dtos.UpdateSchoolDto) (*school_entity.School, error) {
	schoolFound, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	schoolFound.Name = update.Name
	schoolFound.Code = update.Code
	schoolFound.Address = update.Address
	schoolFound.City = update.City
	schoolFound.State = update.State
	schoolFound.ZipCode = update.ZipCode
	schoolFound.Country = update.Country
	schoolFound.PhoneNumber = update.PhoneNumber
	schoolFound.Email = update.Email
	schoolFound.IsActive = update.IsActive
	schoolFound.Description = update.Description

	school, err := s.repo.Update(id, schoolFound)
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (s *SchoolUseCase) Delete(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
