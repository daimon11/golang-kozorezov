package main

import (
	"fmt"
)

// func Run(w io.Writer) error {
// 	_, err := fmt.Fprintln(w, "homeworks: ok")
// 	return err
// }

// Добейся, чтобы каждое уведомление попадало ровно в один из трёх результатов:
// невалидные данные
// → ошибка
// → отправки нет

// валидное обычное уведомление
// → status: queued

// валидное срочное уведомление
// → status: queued-priority

// Требования:
// • Recipient и Subject являются обязательными полями;
// • если пусты оба поля, первой возвращается ошибка получателя;
// • после ошибки уведомление не форматируется и не отправляется;
// • обычное уведомление получает статус queued;
// • срочное уведомление получает статус queued-priority;
// • все пять подготовленных сценариев продолжают запускаться независимо друг от друга.

// Конкретный способ реализации выбери самостоятельно, используя материал первого модуля.

type Notification struct {
	Recipient string
	Subject   string
	Body      string
	Channel   string
	IsUrgent  bool
}

func (n Notification) Format() string {
	return fmt.Sprintf(
		"to: %s | subject: %s | body: %s | channel: %s | urgent: %v",
		n.Recipient,
		n.Subject,
		n.Body,
		n.Channel,
		n.IsUrgent,
	)
}

func (n Notification) Send() (string, string) {
	message := n.Format()
	status := "queued"
	priorityStatus := "queued-priority"
	if n.IsUrgent {
		return message, priorityStatus
	}
	return message, status
}

func (n Notification) Validate() error {
	if n.Recipient == "" {
		return fmt.Errorf("recipient is required")
	}
	if n.Subject == "" {
		return fmt.Errorf("subject is required")
	}

	return nil
}

func runCase(name string, n Notification) {
	fmt.Println("CASE:", name)

	err := n.Validate()
	if err != nil {
		fmt.Println("error:", err)
		fmt.Println("sent: no")
		fmt.Println()
		return
	}

	message, status := n.Send()
	fmt.Println(message)
	fmt.Println("status:", status)
	fmt.Println("sent: yes")
	fmt.Println()
}

func main() {
	regular := Notification{
		Recipient: "user@example.com",
		Subject:   "Deploy done",
		Body:      "v1.2.3 deployed",
		Channel:   "email",
		IsUrgent:  false,
	}

	urgent := Notification{
		Recipient: "ops@example.com",
		Subject:   "Database unavailable",
		Body:      "connection refused",
		Channel:   "telegram",
		IsUrgent:  true,
	}

	emptyRecipient := Notification{
		Subject:  "Deploy done",
		Body:     "v1.2.3 deployed",
		Channel:  "email",
		IsUrgent: false,
	}

	emptySubject := Notification{
		Recipient: "user@example.com",
		Body:      "v1.2.3 deployed",
		Channel:   "email",
		IsUrgent:  true,
	}

	bothEmpty := Notification{
		Body:     "v1.2.3 deployed",
		Channel:  "email",
		IsUrgent: true,
	}

	runCase("regular", regular)
	runCase("urgent", urgent)
	runCase("empty recipient", emptyRecipient)
	runCase("empty subject", emptySubject)
	runCase("both empty", bothEmpty)
}
