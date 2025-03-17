package violate

import (
	"fmt"
	"testing"
)

// Interface besar dengan banyak metode
type Worker interface {
	Work()
	Eat()
	Sleep()
}

// Manusia bisa bekerja, makan, dan tidur
type Human struct{}

func (h Human) Work() {
	fmt.Println("Manusia sedang bekerja")
}

func (h Human) Eat() {
	fmt.Println("Manusia sedang makan")
}

func (h Human) Sleep() {
	fmt.Println("Manusia sedang tidur")
}

// Robot hanya bisa bekerja, tapi harus mengimplementasikan semua metode
type Robot struct{}

func (r Robot) Work() {
	fmt.Println("Robot sedang bekerja")
}

func (r Robot) Eat() {
	panic("Robot tidak bisa makan!") // ❌ Melanggar ISP!
}

func (r Robot) Sleep() {
	panic("Robot tidak bisa tidur!") // ❌ Melanggar ISP!
}

func TestViolateISP(t *testing.T) {
	var w Worker = Robot{} // ❌ Robot dipaksa mengimplementasikan Eat() & Sleep()
	w.Work()
	w.Eat()   // ❌ Error saat dijalankan!
	w.Sleep() // ❌ Error saat dijalankan!
}
