package model

import "fmt"

type SmsNotification struct{}

func (SmsNotification) Send(message string) {
	fmt.Printf("Send Notification via Sms with message: %s\n", message)
}
