package usecase

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"time"

	"github.com/winartodev/attencande-system/config"
	"github.com/winartodev/attencande-system/repository"
	"gopkg.in/gomail.v2"
)

type ReminderUsecaseItf interface {
	Reminder(ctx context.Context) (err error)
}

type ReminderUsecase struct {
	Config             config.Config
	Mailer             *gomail.Message
	EmployeeRepository repository.EmployeeRepositoryItf
	AttendanceUsecase  repository.AttendanceRepositoryItf
}

func NewReminderUsecase(reminderUsecase ReminderUsecase) ReminderUsecaseItf {
	return &ReminderUsecase{
		Config:             reminderUsecase.Config,
		Mailer:             reminderUsecase.Mailer,
		EmployeeRepository: reminderUsecase.EmployeeRepository,
		AttendanceUsecase:  reminderUsecase.AttendanceUsecase,
	}
}

func (ru *ReminderUsecase) Reminder(ctx context.Context) (err error) {
	attendances, err := ru.AttendanceUsecase.GetAllAttendanceDB(ctx)
	if err != nil {
		return err
	}

	for _, attendance := range attendances {
		if time.Now().Unix() == attendance.ClockedIn.Unix() || time.Now().Unix() == attendance.ClockedOut.Unix() {
			employee, err := ru.EmployeeRepository.GetEmployeeByIDDB(ctx, attendance.EmployeeID)
			if err != nil {
				return err
			}

			if time.Now().Unix() == attendance.ClockedIn.Unix() {
				ru.sendToReceiver(employee.Email, "CLOCKED IN", fmt.Sprintf("Reminder Clocked In at : %v", attendance.ClockedOut))
			} else if time.Now().Unix() == attendance.ClockedOut.Unix() {
				ru.sendToReceiver(employee.Email, "CLOCKED OUT", fmt.Sprintf("Reminder Clocked Out at %v", attendance.ClockedOut))
			}
		}
	}

	return nil
}

func (ru *ReminderUsecase) sendToReceiver(receiverEmail string, subject string, body string) {
	auth := smtp.PlainAuth(
		"",
		ru.Config.SMTP.SenderEmail,
		ru.Config.SMTP.Password,
		ru.Config.SMTP.Host,
	)

	msg := "Subject: " + subject + "\n\n" + body

	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", ru.Config.SMTP.Host, ru.Config.SMTP.Port),
		auth,
		ru.Config.SMTP.SenderEmail,
		[]string{receiverEmail},
		[]byte(msg),
	)
	if err != nil {
		log.Fatal(err)
	}
}
