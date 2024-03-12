package service

import (
	"errors"
	"factory-method/model"
)

type NotificationService struct {
	notifier model.Notifier
}

func CreateNotificationService(typeNoti string) (service *NotificationService, err error) {
	switch typeNoti {
	case "email":
		return &NotificationService{
			notifier: model.EmailNotification{},
		}, nil
	case "sms":
		return &NotificationService{
			notifier: model.SmsNotification{},
		}, nil
	}
	return nil, errors.New("type invalid")
}

func (s NotificationService) SendNotification(message string) {
	s.notifier.Send(message)
}
