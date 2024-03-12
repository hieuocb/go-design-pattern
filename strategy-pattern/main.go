package main

import (
	"strategy-pattern/model"
	"strategy-pattern/service"
)

func main() {

	emailService := service.NewNotificationService(
		model.EmailNotification{},
	)

	emailService.SendNotification("Hello")

	smsService := service.NewNotificationService(
		model.SmsNotification{},
	)
	smsService.SendNotification("Hi")
}
