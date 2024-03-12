package service

import (
	"strategy-pattern/model"
)

type NotificationService struct {
	notifier model.Notifier
}

func NewNotificationService(notifier model.Notifier) *NotificationService {
	svc := &NotificationService{
		notifier: notifier,
	}
	return svc
}

func (s *NotificationService) SendNotification(message string) {
	s.notifier.Send(message)
}
