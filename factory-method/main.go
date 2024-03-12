package main

import (
	"factory-method/service"
	"fmt"
)

func main() {
	if notiService, err := service.CreateNotificationService("email"); err != nil {
		fmt.Println(err)
	} else {
		notiService.SendNotification("Hello")
	}

	if notiService, err := service.CreateNotificationService("sms"); err != nil {
		fmt.Println(err)
	} else {
		notiService.SendNotification("hi")
	}
}
