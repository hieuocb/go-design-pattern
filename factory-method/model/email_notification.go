package model

import "fmt"

type EmailNotification struct{}

func (EmailNotification) Send(message string) {
	fmt.Printf("Send Notification via email with message %s\n", message)
}
