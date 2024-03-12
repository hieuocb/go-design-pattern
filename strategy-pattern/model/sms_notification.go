package model

import "fmt"

type SmsNotification struct{}

func (SmsNotification) Send(message string) {
	fmt.Printf("\nSend Ssm Notification message: %s", message)
}
