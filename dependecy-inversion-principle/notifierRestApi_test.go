package dependecyinversionprinciple

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
)

// Notifier interface (DIP Abstraction)
type Notifier interface {
	Send(to, message string) string
}

// EmailSender implementation
type EmailSender struct{}

func (e EmailSender) Send(to, message string) string {
	return fmt.Sprintf("📧 Email sent to %s: %s", to, message)
}

// SMSSender implementation
type SMSSender struct{}

func (s SMSSender) Send(to, message string) string {
	return fmt.Sprintf("📱 SMS sent to %s: %s", to, message)
}

// NotificationService depends on Notifier (Abstraction)
type NotificationService struct {
	notifier Notifier
}

func (n NotificationService) Notify(to, message string) string {
	return n.notifier.Send(to, message)
}

// REST API handler
func SendNotification(c echo.Context) error {
	to := c.QueryParam("to")
	message := c.QueryParam("message")
	method := c.QueryParam("method")

	var notifier Notifier

	// Select notification method
	if method == "sms" {
		notifier = SMSSender{}
	} else {
		notifier = EmailSender{} // Default to Email
	}

	service := NotificationService{notifier: notifier}
	response := service.Notify(to, message)

	return c.JSON(http.StatusOK, map[string]string{"message": response})
}

func TestNotifier(t *testing.T) {
	e := echo.New()
	e.GET("/send", SendNotification)

	// Start server in a separate goroutine
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down server")
		}
	}()

	// Capture OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // Capture CTRL+C

	<-quit
	fmt.Println("\nReceived shutdown signal. Waiting 2 seconds before shutdown...")

	// Wait 2 seconds before shutting down
	time.Sleep(2 * time.Second)

	// Graceful shutdown with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("Server forced to shutdown:", err)
	}

	fmt.Println("Server stopped gracefully after 2 seconds delay")
}
