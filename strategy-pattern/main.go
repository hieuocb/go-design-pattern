package main

func main() {
	emailService := NotificationService{
		notifier: EmailNotification{},
	}
	emailService.SendNotification("Hello")

	smsService := NotificationService{
		notifier: SmsNotification{},
	}
	smsService.SendNotification("Hi")
}
