package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct {}

func (e EmailNotifier) Send(message string) {
	fmt.Printf("Sending message: #{message} (Sender: Email)\n")
}

type SMSNotifier struct {}

func (sms SMSNotifier) Send(message string) {
	fmt.Printf("Sending message: #{message} (Sender: SMS)\n")
}

type TelegramNotifier struct {}

func (tl TelegramNotifier) Send(message string) {
	fmt.Printf("Sending message: #{message} (Sender: Telegram)\n")
}

type NotifierDecorator struct {
	core     *NotifierDecorator
	notifier Notifier
}

func (nd NotifierDecorator) Send(message string) {
	nd.notifier.Send(message)

	if nd.notifier != nil {
		nd.core.Send(message)
	}
}

func (nd NotifierDecorator) Decorate(notifier Notifier) NotifierDecorator {
	return NotifierDecorator{
		core:     &nd,
		notifier: notifier,
	}
}

func NewNotifierDecorator(notifier Notifier) NotifierDecorator {
	return NotifierDecorator{
		notifier: notifier,
	}
}

type Service struct {
	notifier Notifier
}

func (s Service) SendNotification(message string) {
	s.notifier.Send(message)
}

func main() {
	notifier := NewNotifierDecorator(EmailNotifier{}).Decorate(SMSNotifier{}).Decorate(TelegramNotifier{})

	s := Service{notifier: notifier}

	s.SendNotification("Hello, World!")
}