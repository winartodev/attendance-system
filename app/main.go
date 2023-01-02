package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/winartodev/attencande-system/config"
	"github.com/winartodev/attencande-system/repository"
	"github.com/winartodev/attencande-system/server"
	"github.com/winartodev/attencande-system/usecase"
	"gopkg.in/gomail.v2"
)

func main() {
	c := config.Config{}
	cfg := c.NewConfig()

	client, err := config.NewDatabase(cfg)
	if err != nil {
		panic(err)
	}

	err = config.AutoMigrate(client)
	if err != nil {
		log.Fatal("Automigration failed error: %v", err)
	}

	mailer := gomail.NewMessage()

	defer client.Close()

	employeeRepository := repository.NewEmployeeRepository(repository.EmployeeRepository{Client: client})
	employeeUsecase := usecase.NewEmployeeUsecase(usecase.EmployeeUsecase{EmployeeRepository: employeeRepository})
	employeeHandler := server.NewEmployeeUsecase(server.EmployeeHandler{EmployeeUsecase: employeeUsecase})

	attendanceRepository := repository.NewAttendanceRepository(repository.AttendanceRepository{Client: client})
	attendanceUsecase := usecase.NewAttendanceUsecase(usecase.AttendanceUsecase{AttendanceRepository: attendanceRepository, EmployeeRepository: employeeRepository})
	attendanceHandler := server.NewAttendanceUsecase(server.AttendanceHandler{AttendanceUsecase: attendanceUsecase})

	userRepository := repository.NewUserRepository(repository.UserRepository{Client: client})
	userUsecase := usecase.NewUserUsecase(usecase.UserUsecase{UserRepository: userRepository})
	userHandler := server.NewUserUsecase(server.UserHandler{UserUsecase: userUsecase})

	reminderUsecase := usecase.NewReminderUsecase(usecase.ReminderUsecase{Mailer: mailer, Config: cfg, EmployeeRepository: employeeRepository, AttendanceUsecase: attendanceRepository})
	reminderHandler := server.NewReminderHandler(server.ReminderHandler{ReminderUsecase: reminderUsecase})

	s := server.Server{
		Echo:              echo.New(),
		EmployeeHandler:   employeeHandler,
		AttendanceHandler: attendanceHandler,
		UserHandler:       userHandler,
		ReminderHandler:   reminderHandler,
	}

	s.Routes()

	s.StartCron()

	if err := s.Echo.Start(fmt.Sprintf(":%s", cfg.App.Port)); err != nil {
		log.Fatalf("Echo can't start error: %v", err)
	}
}
