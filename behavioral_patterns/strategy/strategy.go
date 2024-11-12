package main

import "fmt"

type Notifier interface {
	Send()
}

type EmailNotifier struct{
	name	 	string
	message 	string
}

func (e EmailNotifier) Send() {
	fmt.Printf("Sending %s notification: %s\n", e.name, e.message)
}

type SMSNotifier struct{
	name	 	string
	message 	string
}

func (sms SMSNotifier) Send() {
	fmt.Printf("Sending %s notification: %s\n", sms.name, sms.message)
}

type NotificationService struct {
	Notifier Notifier
}

func (s NotificationService) SendNotification() {
	s.Notifier.Send()
}


func main() {
	email := NotificationService{Notifier:EmailNotifier{name:"Email",message: "Email notification"}}
	email.SendNotification()

	sms := NotificationService{Notifier:SMSNotifier{name:"SMS" ,message: "SMS notification"}}
	sms.SendNotification()
}