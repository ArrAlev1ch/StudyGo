package main

import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct{}
func (e EmailNotifier) Notify(message string) {
	fmt.Println("Отправлено письмо:", message)
	
}
type TelegramNotifier struct{}
func (t TelegramNotifier) Notify(message string){
	fmt.Println("Отправлено сообщение в Telegram:", message)
}

func SendNotification(n Notifier, msg string) {
	n.Notify(msg)
}

func main() {
	email := EmailNotifier{}
	telegram := TelegramNotifier{}

	SendNotification(telegram, "Сообщение тебе по тгшке")
	SendNotification(email, "Сообщение тебе по почте")
}