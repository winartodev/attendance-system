package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/winartodev/attencande-system/ent"
	"github.com/winartodev/attencande-system/helper"
	"github.com/winartodev/attencande-system/usecase"
)

type AttendanceHandler struct {
	AttendanceUsecase usecase.AttendanceUsecaseItf
	Reminder          ReminderHandler
}

func NewAttendanceUsecase(attendanceHandler AttendanceHandler) AttendanceHandler {
	return AttendanceHandler{
		AttendanceUsecase: attendanceHandler.AttendanceUsecase,
		Reminder:          attendanceHandler.Reminder,
	}
}

// CreateAttendanceHandler handle create new attendance
func (ah *AttendanceHandler) CreateAttendanceHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var attendance ent.Attendance
	err := c.Bind(&attendance)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusUnprocessableEntity),
			Message: err.Error(),
		})
	}

	err = ah.AttendanceUsecase.CreateAttendance(ctx, attendance)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	// after create attendance success will publish reminder
	err = ah.Reminder.PublishReminder(ctx, "reminder.event.publish", attendance)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccess{
		Status:  http.StatusText(http.StatusOK),
		Message: "Attendance success created",
	})

}

// GetAllAttendanceHandler handle get all attendacne
func (ah *AttendanceHandler) GetAllAttendanceHandler(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := ah.AttendanceUsecase.GetAllAttendance(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccess{
		Status:  http.StatusText(http.StatusOK),
		Message: "",
		Data: struct {
			Attendances []*ent.Attendance `json:"attendances"`
		}{Attendances: result},
	})
}

// GetAttendanceByIDHandler handle get attendance by specific id
func (ah *AttendanceHandler) GetAttendanceByIDHandler(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 16, 64)
	ctx := c.Request().Context()

	result, err := ah.AttendanceUsecase.GetAttendanceByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccess{
		Status:  http.StatusText(http.StatusOK),
		Message: "",
		Data: struct {
			Attendance *ent.Attendance `json:"attendance"`
		}{Attendance: result},
	})
}

// UpdateAttendanceByIDHandler will update attendance
func (ah *AttendanceHandler) UpdateAttendanceByIDHandler(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 16, 64)
	ctx := c.Request().Context()

	var attendance ent.Attendance
	err := c.Bind(&attendance)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusUnprocessableEntity),
			Message: err.Error(),
		})
	}

	err = ah.AttendanceUsecase.UpdateAttendance(ctx, id, attendance)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}
	// after update success wil publish reminder
	err = ah.Reminder.PublishReminder(ctx, "reminder.event.publish", attendance)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccess{
		Status:  http.StatusText(http.StatusOK),
		Message: "Attendance success updated",
	})
}

// DeleteAttendanceByIDHandler will delete attendance by specific id
func (ah *AttendanceHandler) DeleteAttendanceByIDHandler(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 16, 64)
	ctx := c.Request().Context()

	err := ah.AttendanceUsecase.DeleteAttendance(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccess{
		Status:  http.StatusText(http.StatusOK),
		Message: "Attendance success deleted",
	})
}
