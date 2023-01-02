package server

import (
	"context"

	"github.com/winartodev/attencande-system/usecase"
)

type ReminderHandler struct {
	ReminderUsecase usecase.ReminderUsecaseItf
}

func NewReminderHandler(employeeHandler ReminderHandler) ReminderHandler {
	return ReminderHandler{
		ReminderUsecase: employeeHandler.ReminderUsecase,
	}
}

func (rh *ReminderHandler) StartReminder(ctx context.Context) (err error) {
	if err := rh.ReminderUsecase.Reminder(ctx); err != nil {
		return err
	}
	return err
}
