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
}

func NewAttendanceUsecase(attendanceHandler AttendanceHandler) AttendanceHandler {
	return AttendanceHandler{
		AttendanceUsecase: attendanceHandler.AttendanceUsecase,
	}
}

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

	return c.JSON(http.StatusOK, helper.ResponseSuccess{
		Status:  http.StatusText(http.StatusOK),
		Message: "Attendance success created",
	})

}

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

	return c.JSON(http.StatusOK, helper.ResponseSuccess{
		Status:  http.StatusText(http.StatusOK),
		Message: "Attendance success updated",
	})
}

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
