package liskovsubstitutionprinciple

import (
	"fmt"
	"testing"
)

// Interface umum Payment
type Payment interface {
	Pay(amount float64) string
}

// Interface Refundable hanya untuk yang bisa refund
type Refundable interface {
	Refund(amount float64) string
}

// CreditCard mendukung Payment & Refundable
type CreditCard struct{}

func (cc CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Dibayar %.2f dengan Kartu Kredit", amount)
}

func (cc CreditCard) Refund(amount float64) string {
	return fmt.Sprintf("Refund %.2f ke Kartu Kredit", amount)
}

// Crypto hanya mendukung Payment
type Crypto struct{}

func (c Crypto) Pay(amount float64) string {
	return fmt.Sprintf("Dibayar %.2f dengan Cryptocurrency", amount)
}

// Fungsi untuk menangani refund hanya jika metode mendukung refund
func ProcessRefund(p Payment, amount float64) string {
	if r, ok := p.(Refundable); ok {
		return r.Refund(amount)
	}
	return "Refund tidak didukung oleh metode pembayaran ini"
}

func TestPayment(t *testing.T) {
	creditCard := CreditCard{}
	crypto := Crypto{}

	fmt.Println(creditCard.Pay(100))
	fmt.Println(ProcessRefund(creditCard, 50)) // ✅ Berhasil refund

	fmt.Println(crypto.Pay(100))
	fmt.Println(ProcessRefund(crypto, 50)) // ✅ Aman (mengembalikan pesan error)
}
