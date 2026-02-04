package repositories

import (
	"context"

	"github.com/HendraaaIrwn/hr-restapi/hr-api/internal/domain/models"

	// models "github.com/HendraaaIrwn/hr-restapi/hr-api/internal/models"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindAll(ctx context.Context) ([]models.Employee, error)
	FindByID(ctx context.Context, id uint) (*models.Employee, error)
	Create(ctx context.Context, employee *models.Employee) error
	Update(ctx context.Context, employee *models.Employee) error
	Delete(ctx context.Context, id uint) error
	SearchByName(ctx context.Context, name string) ([]models.Employee, error)
}

type employeeRepository struct {
	DB *gorm.DB
}

func NewEmployeeRepository(dB *gorm.DB) EmployeeRepository {
	return &employeeRepository{
		DB: dB,
	}
}

// Create implements [EmployeeRepository].
func (r *employeeRepository) Create(ctx context.Context, employee *models.Employee) error {
	return r.DB.WithContext(ctx).Create(employee).Error
}

// Delete implements [EmployeeRepository].
func (r *employeeRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&models.Employee{}, id).Error
}

// FindAll implements [EmployeeRepository].
func (r *employeeRepository) FindAll(ctx context.Context) ([]models.Employee, error) {
	var employees []models.Employee
	err := r.DB.WithContext(ctx).Find(&employees).Error
	return employees, err

}

// FindByID implements [EmployeeRepository].
func (r *employeeRepository) FindByID(ctx context.Context, id uint) (*models.Employee, error) {
	var employee models.Employee
	err := r.DB.WithContext(ctx).First(&employee, id).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

// Update implements [EmployeeRepository].
func (r *employeeRepository) Update(ctx context.Context, employee *models.Employee) error {
	return r.DB.WithContext(ctx).Save(employee).Error
}

// SearchByName implements [EmployeeRepository].
func (r *employeeRepository) SearchByName(ctx context.Context, name string) ([]models.Employee, error) {
	var employees []models.Employee
	pattern := "%" + name + "%"
	err := r.DB.WithContext(ctx).
		Where("first_name ILIKE ? OR last_name ILIKE ?", pattern, pattern).
		Find(&employees).Error
	return employees, err
}
