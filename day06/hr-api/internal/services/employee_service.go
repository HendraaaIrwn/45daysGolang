package services

import (
	"context"
	"time"

	"github.com/HendraaaIrwn/hr-restapi/hr-api/internal/domain/models"
	"github.com/HendraaaIrwn/hr-restapi/hr-api/internal/dto"
	errs "github.com/HendraaaIrwn/hr-restapi/hr-api/internal/errors"
	"github.com/HendraaaIrwn/hr-restapi/hr-api/internal/repositories"
	"github.com/go-playground/validator/v10"
)

const employeeDateLayout = "2006-01-02"

// EmployeeServiceInterface defines the service interface
type EmployeeServiceInterface interface {
	Create(ctx context.Context, req *dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error)
	FindByID(ctx context.Context, id uint) (*dto.EmployeeResponse, error)
	GetAll(ctx context.Context) ([]dto.EmployeeResponse, error)
	Update(ctx context.Context, id uint, req *dto.UpdateEmployeeRequest) (*dto.EmployeeResponse, error)
	Delete(ctx context.Context, id uint) error
	SearchByName(ctx context.Context, name string) ([]dto.EmployeeResponse, error)
}

type employeeService struct {
	repo     repositories.EmployeeRepository
	validate *validator.Validate
}

func NewEmployeeService(repo repositories.EmployeeRepository) EmployeeServiceInterface {
	return &employeeService{
		repo:     repo,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

// Create creates a new employee
func (s *employeeService) Create(ctx context.Context, req *dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, errs.ErrInvalidInput
	}

	hireDate, err := time.Parse(employeeDateLayout, req.HireDate)
	if err != nil {
		return nil, errs.ErrInvalidInput
	}

	firstName := req.FirstName
	employee := &models.Employee{
		FirstName: &firstName,
		LastName:  req.LastName,
		Email:     req.Email,
		HireDate:  hireDate,
		JobID:     int32(req.JobID),
	}

	if err := s.repo.Create(ctx, employee); err != nil {
		return nil, err
	}

	resp := employeeResponseFromModel(employee)
	return &resp, nil
}

// FindByID gets an employee by ID
func (s *employeeService) FindByID(ctx context.Context, id uint) (*dto.EmployeeResponse, error) {
	emp, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := employeeResponseFromModel(emp)
	return &resp, nil
}

// GetAll gets all employees
func (s *employeeService) GetAll(ctx context.Context) ([]dto.EmployeeResponse, error) {
	emps, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []dto.EmployeeResponse
	for i := range emps {
		responses = append(responses, employeeResponseFromModel(&emps[i]))
	}

	return responses, nil
}

// Update updates an employee
func (s *employeeService) Update(ctx context.Context, id uint, req *dto.UpdateEmployeeRequest) (*dto.EmployeeResponse, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, errs.ErrInvalidInput
	}

	emp, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.FirstName != nil {
		emp.FirstName = req.FirstName
	}
	if req.LastName != nil {
		emp.LastName = *req.LastName
	}
	if req.Email != nil {
		emp.Email = *req.Email
	}
	if req.HireDate != nil {
		hireDate, err := time.Parse(employeeDateLayout, *req.HireDate)
		if err != nil {
			return nil, errs.ErrInvalidInput
		}
		emp.HireDate = hireDate
	}
	if req.JobID != nil {
		emp.JobID = int32(*req.JobID)
	}

	if err := s.repo.Update(ctx, emp); err != nil {
		return nil, err
	}

	resp := employeeResponseFromModel(emp)
	return &resp, nil
}

// Delete deletes an employee by ID
func (s *employeeService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

// SearchByName searches employees by name
func (s *employeeService) SearchByName(ctx context.Context, name string) ([]dto.EmployeeResponse, error) {
	if name == "" {
		return s.GetAll(ctx)
	}

	emps, err := s.repo.SearchByName(ctx, name)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.EmployeeResponse, 0, len(emps))
	for i := range emps {
		responses = append(responses, employeeResponseFromModel(&emps[i]))
	}
	return responses, nil
}

func employeeResponseFromModel(emp *models.Employee) dto.EmployeeResponse {
	firstName := ""
	if emp.FirstName != nil {
		firstName = *emp.FirstName
	}
	return dto.EmployeeResponse{
		EmployeeID: uint(emp.EmployeeID),
		FirstName:  firstName,
		LastName:   emp.LastName,
		Email:      emp.Email,
		HireDate:   emp.HireDate.Format(employeeDateLayout),
		JobID:      uint(emp.JobID),
	}
}
