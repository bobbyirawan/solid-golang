package openclosedprinciple

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
)

// 1️⃣ Define the PaymentMethod interface
type PaymentMethod interface {
	Pay(amount float64) string
}

// 2️⃣ Implement Credit Card payment
type CreditCard struct{}

func (cc CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Processed credit card payment of %.2f", amount)
}

// 3️⃣ Implement PayPal payment
type PayPal struct{}

func (pp PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Processed PayPal payment of %.2f", amount)
}

// 4️⃣ Implement Bank Transfer payment
type BankTransfer struct{}

func (bt BankTransfer) Pay(amount float64) string {
	return fmt.Sprintf("Processed bank transfer payment of %.2f", amount)
}

// 5️⃣ Implement Crypto payment
type Crypto struct{}

func (c Crypto) Pay(amount float64) string {
	return fmt.Sprintf("Processed cryptocurrency payment of %.2f", amount)
}

// 6️⃣ PaymentProcessor struct to handle payments dynamically
type PaymentProcessor struct {
	methods map[string]PaymentMethod
}

// 7️⃣ Constructor function for PaymentProcessor
func NewPaymentProcessor() *PaymentProcessor {
	return &PaymentProcessor{
		methods: map[string]PaymentMethod{
			"credit_card":   CreditCard{},
			"paypal":        PayPal{},
			"bank_transfer": BankTransfer{},
			"crypto":        Crypto{},
		},
	}
}

// 8️⃣ Process payment dynamically based on user input
func (pp *PaymentProcessor) ProcessPayment(method string, amount float64) string {
	paymentMethod, exists := pp.methods[method]
	if !exists {
		return "Invalid payment method"
	}
	return paymentMethod.Pay(amount)
}

// 9️⃣ Setup REST API with `echo`
func TestRestMultiPay(t *testing.T) {
	e := echo.New()
	paymentProcessor := NewPaymentProcessor()

	// Route to process payments
	e.POST("/pay", func(c echo.Context) error {
		var request struct {
			Method string  `json:"method"`
			Amount float64 `json:"amount"`
		}

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		result := paymentProcessor.ProcessPayment(request.Method, request.Amount)
		return c.JSON(http.StatusOK, map[string]string{"message": result})
	})

	// Start the server
	e.Logger.Fatal(e.Start(":8080")) // API runs on localhost:8080
}
