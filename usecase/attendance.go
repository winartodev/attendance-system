package usecase

import (
	"context"

	"github.com/winartodev/attencande-system/ent"
	"github.com/winartodev/attencande-system/repository"
)

type AttendanceUsecaseItf interface {
	GetAllAttendance(ctx context.Context) (result []*ent.Attendance, err error)
	GetAttendanceByID(ctx context.Context, id int64) (result *ent.Attendance, err error)
	CreateAttendance(ctx context.Context, attendance ent.Attendance) (err error)
	UpdateAttendance(ctx context.Context, id int64, attendance ent.Attendance) (err error)
	DeleteAttendance(ctx context.Context, id int64) (err error)
}

type AttendanceUsecase struct {
	EmployeeRepository   repository.EmployeeRepositoryItf
	AttendanceRepository repository.AttendanceRepositoryItf
}

func NewAttendanceUsecase(attendanceUsecase AttendanceUsecase) AttendanceUsecaseItf {
	return &AttendanceUsecase{
		AttendanceRepository: attendanceUsecase.AttendanceRepository,
		EmployeeRepository:   attendanceUsecase.EmployeeRepository,
	}
}

func (au *AttendanceUsecase) GetAllAttendance(ctx context.Context) (result []*ent.Attendance, err error) {
	result, err = au.AttendanceRepository.GetAllAttendanceDB(ctx)
	if err != nil {
		return result, err
	}

	return result, err
}

func (au *AttendanceUsecase) GetAttendanceByID(ctx context.Context, id int64) (result *ent.Attendance, err error) {
	result, err = au.AttendanceRepository.GetAttendanceByIDDB(ctx, id)
	if err != nil {
		return result, err
	}

	return result, err
}

func (au *AttendanceUsecase) CreateAttendance(ctx context.Context, attendance ent.Attendance) (err error) {
	_, err = au.EmployeeRepository.GetEmployeeByIDDB(ctx, attendance.EmployeeID)
	if err != nil {
		return err
	}

	err = au.AttendanceRepository.CreateAttendanceDB(ctx, attendance)
	if err != nil {
		return err
	}

	return err
}

func (au *AttendanceUsecase) UpdateAttendance(ctx context.Context, id int64, attendance ent.Attendance) (err error) {
	err = au.AttendanceRepository.UpdateAttendanceByIDDB(ctx, id, attendance)
	if err != nil {
		return err
	}

	return err
}

func (au *AttendanceUsecase) DeleteAttendance(ctx context.Context, id int64) (err error) {
	err = au.AttendanceRepository.DeleteAttendanceByIDDB(ctx, id)
	if err != nil {
		return err
	}

	return err
}
