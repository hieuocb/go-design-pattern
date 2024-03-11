package main

import "fmt"

type EmailNotification struct{}

func (EmailNotification) Send(message string) {
	fmt.Printf("\nSend Email Notification message: %s", message)
}
