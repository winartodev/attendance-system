package repository

import (
	"context"

	"github.com/winartodev/attencande-system/ent"
	"github.com/winartodev/attencande-system/ent/employee"
)

type Notification struct {
	Email      string `json:"email"`
	ClockedIn  string `json:"clocked_id"`
	ClockedOut string `json:"clocked_out"`
}

type EmployeeRepositoryItf interface {
	CreateEmployeeDB(ctx context.Context, employee ent.Employee) (err error)
	GetAllEmployeeDB(ctx context.Context) (result []*ent.Employee, err error)
	GetEmployeeByIDDB(ctx context.Context, id int64) (result *ent.Employee, err error)
	UpdateEmployeeByIDDB(ctx context.Context, id int64, employee ent.Employee) (err error)
	DeleteEmployeeByIDDB(ctx context.Context, id int64) (err error)
	Notification(ctx context.Context) (err error)
}

type EmployeeRepository struct {
	Client *ent.Client
}

func NewEmployeeRepository(employeeRepository EmployeeRepository) EmployeeRepositoryItf {
	return &EmployeeRepository{
		Client: employeeRepository.Client,
	}
}

func (er *EmployeeRepository) CreateEmployeeDB(ctx context.Context, employee ent.Employee) (err error) {
	_, err = er.Client.Employee.
		Create().
		SetName(employee.Name).
		SetEmail(employee.Email).
		Save(ctx)

	if err != nil {
		return err
	}

	return err
}

func (er *EmployeeRepository) GetAllEmployeeDB(ctx context.Context) (result []*ent.Employee, err error) {
	result, err = er.Client.Employee.
		Query().
		Select(employee.FieldID, employee.FieldName, employee.FieldEmail).
		All(ctx)
	if err != nil {
		return result, err
	}

	return result, err
}

func (er *EmployeeRepository) GetEmployeeByIDDB(ctx context.Context, id int64) (result *ent.Employee, err error) {
	result, err = er.Client.Employee.
		Query().
		Select(employee.FieldID, employee.FieldName, employee.FieldEmail).
		Where(employee.ID(id)).
		Only(ctx)
	if err != nil {
		return result, err
	}

	return result, err
}

func (er *EmployeeRepository) UpdateEmployeeByIDDB(ctx context.Context, id int64, employee ent.Employee) (err error) {
	_, err = er.Client.Employee.
		UpdateOneID(id).
		SetName(employee.Name).
		SetEmail(employee.Email).
		Save(ctx)
	if err != nil {
		return err
	}

	return err
}

func (er *EmployeeRepository) DeleteEmployeeByIDDB(ctx context.Context, id int64) (err error) {
	err = er.Client.Employee.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return err
	}

	return err
}

func (er *EmployeeRepository) Notification(ctx context.Context) (err error) {
	return err
}
