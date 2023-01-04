package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/winartodev/attencande-system/ent"
	"github.com/winartodev/attencande-system/helper"
	"github.com/winartodev/attencande-system/usecase"
)

type EmployeeHandler struct {
	EmployeeUsecase usecase.EmployeeUsecaseItf
}

func NewEmployeeUsecase(employeeHandler EmployeeHandler) EmployeeHandler {
	return EmployeeHandler{
		EmployeeUsecase: employeeHandler.EmployeeUsecase,
	}
}

func (eh *EmployeeHandler) CreateEmployeeHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var employee ent.Employee
	err := c.Bind(&employee)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusUnprocessableEntity),
			Message: err.Error(),
		})
	}

	err = eh.EmployeeUsecase.CreateEmployee(ctx, employee)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccess{
		Status:  http.StatusText(http.StatusOK),
		Message: "Employee success created",
	})
}

func (eh *EmployeeHandler) GetAllEmployeeHandler(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := eh.EmployeeUsecase.GetAllEmployee(ctx)
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
			Employee []*ent.Employee `json:"employees"`
		}{Employee: result},
	})
}

func (eh *EmployeeHandler) GetEmployeeByIDHandler(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 16, 64)
	ctx := c.Request().Context()

	result, err := eh.EmployeeUsecase.GetEmployeeByID(ctx, id)
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
			Employee *ent.Employee `json:"employee"`
		}{Employee: result},
	})
}

func (eh *EmployeeHandler) UpdateEmployeeByIDHandler(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 16, 64)
	ctx := c.Request().Context()

	var employee ent.Employee
	err := c.Bind(&employee)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusUnprocessableEntity),
			Message: err.Error(),
		})
	}

	err = eh.EmployeeUsecase.UpdateEmployee(ctx, id, employee)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccess{
		Status:  http.StatusText(http.StatusOK),
		Message: "Employee success updated",
	})
}

func (eh *EmployeeHandler) DeleteEmployeeByIDHandler(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 16, 64)
	ctx := c.Request().Context()

	err := eh.EmployeeUsecase.DeleteEmployee(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccess{
		Status:  http.StatusText(http.StatusOK),
		Message: "Employee success deleted",
	})
}
