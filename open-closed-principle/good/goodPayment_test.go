package openclosedprinciple

import (
	"fmt"
	"testing"
)

// 1️⃣ PaymentMethod interface (Open for Extension)
type PaymentMethod interface {
	Pay(amount float64)
}

// 2️⃣ CreditCard struct implementing PaymentMethod
type CreditCard struct{}

func (cc CreditCard) Pay(amount float64) {
	fmt.Println("Processing credit card payment of", amount)
}

// 3️⃣ PayPal struct implementing PaymentMethod
type PayPal struct{}

func (pp PayPal) Pay(amount float64) {
	fmt.Println("Processing PayPal payment of", amount)
}

// 4️⃣ PaymentProcessor that works with multiple payment methods
type PaymentProcessor struct{}

func (pp PaymentProcessor) ProcessPayment(pm PaymentMethod, amount float64) {
	pm.Pay(amount) // ⬅ Polymorphism happens here!
}

func TestGoodPayment(t *testing.T) {
	processor := PaymentProcessor{}

	// Using different payment methods without modifying existing code
	creditCard := CreditCard{}
	paypal := PayPal{}

	processor.ProcessPayment(creditCard, 100)
	processor.ProcessPayment(paypal, 200)
}
