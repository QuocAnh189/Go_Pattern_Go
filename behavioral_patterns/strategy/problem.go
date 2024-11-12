package main

import "fmt"

type INotificationService struct {
	notifierTye string
}

func (n INotificationService) SendNotification(message string) {
	if n.notifierTye == "email" {
		fmt.Println("Sending email notification:", message)
	} else if n.notifierTye == "sms" {
		fmt.Println("Sending SMS notification:", message)
	} else {
		fmt.Println("Unknown notification type")
	}
}

func problem() {
	service := INotificationService{notifierTye: "email"}
	service.SendNotification("Hello Email")
}