package good

import (
	"fmt"
	"testing"
)

// Interface terpisah untuk setiap tanggung jawab
type Worker interface {
	Work()
}

type Eater interface {
	Eat()
}

type Sleeper interface {
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

// Robot hanya bekerja dan tidak dipaksa mengimplementasikan Eat() & Sleep()
type Robot struct{}

func (r Robot) Work() {
	fmt.Println("Robot sedang bekerja")
}

func TestGoodISP(t *testing.T) {
	// Manusia bisa melakukan semua
	human := Human{}
	human.Work()
	human.Eat()
	human.Sleep()

	// Robot hanya bisa bekerja
	robot := Robot{}
	robot.Work()
}
