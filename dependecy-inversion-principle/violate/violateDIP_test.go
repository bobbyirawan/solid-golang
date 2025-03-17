package violate

import (
	"fmt"
	"testing"
)

// EmailSender mengirim email langsung (tanpa abstraction)
type EmailSender struct{}

func (e EmailSender) SendEmail(to, message string) {
	fmt.Printf("Sending Email to %s: %s\n", to, message)
}

// High-Level Module: NotificationService
type NotificationService struct {
	emailSender EmailSender // ❌ Bergantung langsung pada EmailSender
}

func (n NotificationService) Notify(user string, message string) {
	n.emailSender.SendEmail(user, message) // ❌ Ketergantungan langsung
}

func TestViolateDIP(t *testing.T) {
	service := NotificationService{emailSender: EmailSender{}}
	service.Notify("user@example.com", "Hello, this is a notification!")
}
