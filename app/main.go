package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	defer client.Close()

	err = config.AutoMigrate(client)
	if err != nil {
		log.Fatal("Automigration failed error: %v", err)
	}

	rmq, err := config.NewRabbitMQ(cfg)
	if err != nil {
		log.Fatal("Automigration failed error: %v", err)
	}
	defer rmq.Shutdown()

	mailer := gomail.NewMessage()

	employeeRepository := repository.NewEmployeeRepository(repository.EmployeeRepository{Client: client})
	employeeUsecase := usecase.NewEmployeeUsecase(usecase.EmployeeUsecase{EmployeeRepository: employeeRepository})
	employeeHandler := server.NewEmployeeUsecase(server.EmployeeHandler{EmployeeUsecase: employeeUsecase})

	reminderHandler := server.NewReminderHandler(server.ReminderHandler{RabbitMQ: rmq, Mailer: mailer, Config: cfg, EmployeeUsecase: employeeUsecase})

	attendanceRepository := repository.NewAttendanceRepository(repository.AttendanceRepository{Client: client})
	attendanceUsecase := usecase.NewAttendanceUsecase(usecase.AttendanceUsecase{AttendanceRepository: attendanceRepository, EmployeeRepository: employeeRepository})
	attendanceHandler := server.NewAttendanceUsecase(server.AttendanceHandler{AttendanceUsecase: attendanceUsecase, Reminder: reminderHandler})

	userRepository := repository.NewUserRepository(repository.UserRepository{Client: client})
	userUsecase := usecase.NewUserUsecase(usecase.UserUsecase{UserRepository: userRepository})
	userHandler := server.NewUserUsecase(server.UserHandler{UserUsecase: userUsecase})

	s := server.Server{
		Echo:              echo.New(),
		EmployeeHandler:   employeeHandler,
		AttendanceHandler: attendanceHandler,
		UserHandler:       userHandler,
		ReminderHandler:   reminderHandler,
	}

	s.Echo.Use(middleware.CORS())

	s.Routes()

	reminderHandler.Consumer()

	if err := s.Echo.Start(fmt.Sprintf(":%s", cfg.App.Port)); err != nil {
		log.Fatalf("Echo can't start error: %v", err)
	}
}
