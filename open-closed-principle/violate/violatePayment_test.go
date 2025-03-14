package openclosedprinciple

import (
	"fmt"
	"testing"
)

// PaymentService handles payments using different methods
type PaymentService struct{}

func (ps PaymentService) ProcessPayment(method string, amount float64) {
	if method == "credit_card" {
		fmt.Println("Processing credit card payment of", amount)
	} else if method == "paypal" {
		fmt.Println("Processing PayPal payment of", amount)
	} else {
		fmt.Println("Unknown payment method")
	}
}

func TestPayment(t *testing.T) {
	service := PaymentService{}
	service.ProcessPayment("credit_card", 100)
	service.ProcessPayment("paypal", 200)
}
