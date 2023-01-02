package usecase

import (
	"context"

	"github.com/winartodev/attencande-system/ent"
	"github.com/winartodev/attencande-system/repository"
)

type EmployeeUsecaseItf interface {
	GetAllEmployee(ctx context.Context) (result []*ent.Employee, err error)
	GetEmployeeByID(ctx context.Context, id int64) (result *ent.Employee, err error)
	CreateEmployee(ctx context.Context, employee ent.Employee) (err error)
	UpdateEmployee(ctx context.Context, id int64, employee ent.Employee) (err error)
	DeleteEmployee(ctx context.Context, id int64) (err error)
	Reminder(ctx context.Context) (err error)
}

type EmployeeUsecase struct {
	EmployeeRepository repository.EmployeeRepositoryItf
}

func NewEmployeeUsecase(employeeUsecase EmployeeUsecase) EmployeeUsecaseItf {
	return &EmployeeUsecase{
		EmployeeRepository: employeeUsecase.EmployeeRepository,
	}
}

func (eu *EmployeeUsecase) GetAllEmployee(ctx context.Context) (result []*ent.Employee, err error) {
	result, err = eu.EmployeeRepository.GetAllEmployeeDB(ctx)
	if err != nil {
		return result, err
	}

	return result, err
}

func (eu *EmployeeUsecase) GetEmployeeByID(ctx context.Context, id int64) (result *ent.Employee, err error) {
	result, err = eu.EmployeeRepository.GetEmployeeByIDDB(ctx, id)
	if err != nil {
		return result, err
	}

	return result, err
}

func (eu *EmployeeUsecase) CreateEmployee(ctx context.Context, employee ent.Employee) (err error) {
	err = eu.EmployeeRepository.CreateEmployeeDB(ctx, employee)
	if err != nil {
		return err
	}

	return err
}

func (eu *EmployeeUsecase) UpdateEmployee(ctx context.Context, id int64, employee ent.Employee) (err error) {
	err = eu.EmployeeRepository.UpdateEmployeeByIDDB(ctx, id, employee)
	if err != nil {
		return err
	}

	return err
}

func (eu *EmployeeUsecase) DeleteEmployee(ctx context.Context, id int64) (err error) {
	err = eu.EmployeeRepository.DeleteEmployeeByIDDB(ctx, id)
	if err != nil {
		return err
	}

	return err
}

func (eu *EmployeeUsecase) Reminder(ctx context.Context) (err error) {
	err = eu.EmployeeRepository.Notification(ctx)
	if err != nil {
		return err
	}

	return err
}
