package repository

import (
	"context"

	"github.com/winartodev/attencande-system/ent"
	"github.com/winartodev/attencande-system/ent/attendance"
)

type AttendanceRepositoryItf interface {
	CreateAttendanceDB(ctx context.Context, attendance ent.Attendance) (err error)
	GetAllAttendanceDB(ctx context.Context) (result []*ent.Attendance, err error)
	GetAttendanceByIDDB(ctx context.Context, id int64) (result *ent.Attendance, err error)
	UpdateAttendanceByIDDB(ctx context.Context, id int64, attendance ent.Attendance) (err error)
	DeleteAttendanceByIDDB(ctx context.Context, id int64) (err error)
}

type AttendanceRepository struct {
	Client *ent.Client
}

func NewAttendanceRepository(attendanceRepository AttendanceRepository) AttendanceRepositoryItf {
	return &AttendanceRepository{
		Client: attendanceRepository.Client,
	}
}

func (ar *AttendanceRepository) CreateAttendanceDB(ctx context.Context, attendance ent.Attendance) (err error) {
	_, err = ar.Client.Attendance.
		Create().
		SetEmployeeID(attendance.EmployeeID).
		SetClockedIn(attendance.ClockedIn).
		SetClockedOut(attendance.ClockedOut).
		Save(ctx)

	if err != nil {
		return err
	}

	return err
}

func (ar *AttendanceRepository) GetAllAttendanceDB(ctx context.Context) (result []*ent.Attendance, err error) {
	result, err = ar.Client.Attendance.
		Query().
		Select(attendance.Columns...).
		All(ctx)
	if err != nil {
		return result, err
	}

	return result, err
}

func (ar *AttendanceRepository) GetAttendanceByIDDB(ctx context.Context, id int64) (result *ent.Attendance, err error) {
	result, err = ar.Client.Attendance.
		Query().
		Select(attendance.Columns...).
		Where(attendance.ID(id)).
		Only(ctx)
	if err != nil {
		return result, err
	}

	return result, err
}

func (ar *AttendanceRepository) UpdateAttendanceByIDDB(ctx context.Context, id int64, attendance ent.Attendance) (err error) {
	_, err = ar.Client.Attendance.
		UpdateOneID(id).
		SetClockedIn(attendance.ClockedIn).
		SetClockedOut(attendance.ClockedOut).
		Save(ctx)
	if err != nil {
		return err
	}

	return err
}

func (ar *AttendanceRepository) DeleteAttendanceByIDDB(ctx context.Context, id int64) (err error) {
	err = ar.Client.Attendance.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return err
	}

	return err
}
