package good

import (
	"fmt"
	"testing"
)

// Interface Notifier (abstraksi)
type Notifier interface {
	Send(to, message string)
}

// Implementasi EmailSender (mengikuti interface)
type EmailSender struct{}

func (e EmailSender) Send(to, message string) {
	fmt.Printf("📧 Sending Email to %s: %s\n", to, message)
}

// Implementasi SMS Sender
type SMSSender struct{}

func (s SMSSender) Send(to, message string) {
	fmt.Printf("📱 Sending SMS to %s: %s\n", to, message)
}

// High-Level Module: Tidak bergantung pada implementasi konkret
type NotificationService struct {
	notifier Notifier // ✅ Bergantung pada interface
}

func (n NotificationService) Notify(user string, message string) {
	n.notifier.Send(user, message)
}

func TestGoodDIP(t *testing.T) {
	// Bisa pakai Email
	emailService := NotificationService{notifier: EmailSender{}}
	emailService.Notify("user@example.com", "Hello via Email!")

	// Bisa pakai SMS tanpa mengubah NotificationService
	smsService := NotificationService{notifier: SMSSender{}}
	smsService.Notify("123456789", "Hello via SMS!")
}
