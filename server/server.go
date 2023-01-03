package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo"
	"github.com/winartodev/attencande-system/middleware"
)

type Server struct {
	Echo              *echo.Echo
	EmployeeHandler   EmployeeHandler
	AttendanceHandler AttendanceHandler
	UserHandler       UserHandler
	ReminderHandler   ReminderHandler
}

func healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}

func (s *Server) Routes() {
	admin := s.Echo.Group("/admin")
	admin.Use(middleware.Authentication)
	{
		admin.POST("/employees", s.EmployeeHandler.CreateEmployeeHandler)
		admin.GET("/employees", s.EmployeeHandler.GetAllEmployeeHandler)
		admin.GET("/employees/:id", s.EmployeeHandler.GetEmployeeByIDHandler)
		admin.PUT("/employees/:id", s.EmployeeHandler.UpdateEmployeeByIDHandler)
		admin.DELETE("/employees/:id", s.EmployeeHandler.DeleteEmployeeByIDHandler)

		admin.POST("/attendances", s.AttendanceHandler.CreateAttendanceHandler)
		admin.GET("/attendances", s.AttendanceHandler.GetAllAttendanceHandler)
		admin.GET("/attendances/:id", s.AttendanceHandler.GetAttendanceByIDHandler)
		admin.PUT("/attendances/:id", s.AttendanceHandler.UpdateAttendanceByIDHandler)
		admin.DELETE("/attendances/:id", s.AttendanceHandler.DeleteAttendanceByIDHandler)
	}

	user := s.Echo.Group("/users")
	{
		user.POST("/login", s.UserHandler.LoginHandler)
		user.POST("/register", s.UserHandler.RegisterHandler)
		user.POST("/logout", s.UserHandler.LogoutHandler)
	}

	s.Echo.GET("/healthz", healthz)
}
