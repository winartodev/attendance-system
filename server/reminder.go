package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/winartodev/attencande-system/config"
	"github.com/winartodev/attencande-system/ent"
	"github.com/winartodev/attencande-system/usecase"
	"gopkg.in/gomail.v2"
)

type ReminderPayload struct {
	Attendacne  ent.Attendance `json:"attendance"`
	IsClockedIn bool           `json:"is_clocked_in"`
}

type ReminderHandler struct {
	RabbitMQ        *config.RabbitMQ
	Mailer          *gomail.Message
	Config          config.Config
	EmployeeUsecase usecase.EmployeeUsecaseItf
}

func NewReminderHandler(employeeHandler ReminderHandler) ReminderHandler {
	return ReminderHandler{
		RabbitMQ:        employeeHandler.RabbitMQ,
		Mailer:          employeeHandler.Mailer,
		Config:          employeeHandler.Config,
		EmployeeUsecase: employeeHandler.EmployeeUsecase,
	}
}

// PublishReminder is function to publish message to rabbitmq
func (rh *ReminderHandler) PublishReminder(ctx context.Context, routingKey string, attendance ent.Attendance) error {
	var reminders []ReminderPayload
	// validate clockedin and clockedout
	if !attendance.ClockedIn.IsZero() {
		reminders = append(reminders, ReminderPayload{
			Attendacne:  attendance,
			IsClockedIn: true,
		})
	}

	if !attendance.ClockedOut.IsZero() {
		reminders = append(reminders, ReminderPayload{
			Attendacne:  attendance,
			IsClockedIn: false,
		})
	}

	for _, reminder := range reminders {
		headers := make(amqp.Table)
		if reminder.IsClockedIn {
			headers["x-delay"] = time.Until(reminder.Attendacne.ClockedIn).Milliseconds()
		} else {
			headers["x-delay"] = time.Until(reminder.Attendacne.ClockedOut).Milliseconds()
		}

		body, _ := json.Marshal(reminder)
		err := rh.RabbitMQ.Chann.PublishWithContext(ctx, "delayed", routingKey, false, false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			Timestamp:    time.Now(),
			ContentType:  "application/json",
			Body:         body,
			Headers:      headers,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// Consumer will consume message which has been published
func (rh *ReminderHandler) Consumer() error {
	var err error
	delayedQueue, err := rh.RabbitMQ.Chann.QueueDeclare("reminder-published-queue", true, false, false, false, nil)
	if err != nil {
		return err
	}
	err = rh.RabbitMQ.Chann.QueueBind(delayedQueue.Name, "reminder.event.publish", "delayed", false, nil)
	if err != nil {
		return err
	}

	err = rh.RabbitMQ.Chann.Qos(2, 0, false)
	if err != nil {
		return err
	}

	published, err := rh.RabbitMQ.Chann.Consume(
		"reminder-published-queue",
		"reminder-published-consumer",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		return err
	}
	go rh.consume(published)
	return nil
}

func (rh *ReminderHandler) consume(ds <-chan amqp.Delivery) {
	for {
		select {
		case d, ok := <-ds:
			if !ok {
				return
			}
			rh.evaluateMessage(d.Body)
			d.Ack(false)
		}
	}
}

// evaluateMessage function to evaluate message which has been received by the consumer and send it to receiver
func (rh *ReminderHandler) evaluateMessage(body []byte) {
	var reminder ReminderPayload
	json.Unmarshal(body, &reminder)
	employee, _ := rh.EmployeeUsecase.GetEmployeeByID(context.Background(), reminder.Attendacne.EmployeeID)
	if reminder.IsClockedIn {
		go rh.sendToReceiver(employee.Email, "CLOCKED IN", fmt.Sprintf("Reminder Clocked In at : %v", reminder.Attendacne.ClockedIn))
	} else {
		go rh.sendToReceiver(employee.Email, "CLOCKED OUT", fmt.Sprintf("Reminder Clocked out at : %v", reminder.Attendacne.ClockedOut))
	}
}

// sendToReceiver will send email to specific receiver email
func (rh *ReminderHandler) sendToReceiver(receiverEmail string, subject string, body string) {
	auth := smtp.PlainAuth(
		"",
		os.Getenv("SMTP_SENDER_EMAIL"),
		os.Getenv("SMTP_PASSWORD"),
		rh.Config.SMTP.Host,
	)

	msg := "Subject: " + subject + "\n\n" + body

	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", rh.Config.SMTP.Host, rh.Config.SMTP.Port),
		auth,
		os.Getenv("SMTP_SENDER_EMAIL"),
		[]string{receiverEmail},
		[]byte(msg),
	)
	if err != nil {
		log.Fatal(err)
	}
}
