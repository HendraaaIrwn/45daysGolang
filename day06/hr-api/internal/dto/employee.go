package dto

// CreateEmployeeRequest for creating an employee
type CreateEmployeeRequest struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=100"` // Validation tags
	LastName  string `json:"last_name" validate:"required,min=2,max=100"`
	Email     string `json:"email" validate:"required,email"`
	HireDate  string `json:"hire_date" validate:"required,datetime=2006-01-02"`
	JobID     uint   `json:"job_id" validate:"required"`
}

// UpdateEmployeeRequest for partial update
type UpdateEmployeeRequest struct {
	FirstName *string `json:"first_name,omitempty" validate:"omitempty,min=2,max=100"`
	LastName  *string `json:"last_name,omitempty" validate:"omitempty,min=2,max=100"`
	Email     *string `json:"email,omitempty" validate:"omitempty,email"`
	HireDate  *string `json:"hire_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
	JobID     *uint   `json:"job_id,omitempty" validate:"omitempty,gt=0"`
}

// EmployeeResponse for output (hide internal fields if needed)
type EmployeeResponse struct {
	EmployeeID uint   `json:"employee_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	HireDate   string `json:"hire_date"`
	JobID      uint   `json:"job_id"`
}
