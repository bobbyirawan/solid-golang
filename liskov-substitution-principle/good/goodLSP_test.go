package liskovsubstitutionprinciple

import (
	"fmt"
	"testing"
)

// 1️⃣ Interface for flying birds
type FlyingBird interface {
	Fly()
}

// 2️⃣ Interface for all birds
type Bird interface {
	Eat()
}

// 3️⃣ Sparrow implements both Bird and FlyingBird
type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow is flying!")
}

func (s Sparrow) Eat() {
	fmt.Println("Sparrow is eating!")
}

// 4️⃣ Penguin implements only Bird (but not FlyingBird)
type Penguin struct{}

func (p Penguin) Eat() {
	fmt.Println("Penguin is eating!")
}

// 5️⃣ Function that correctly handles only FlyingBirds
func MakeBirdFly(fb FlyingBird) {
	fb.Fly()
}

func TestFlyingBirds(t *testing.T) {
	sparrow := Sparrow{}
	penguin := Penguin{}

	MakeBirdFly(sparrow) // ✅ Works fine
	// MakeBirdFly(penguin) ❌ Compiler error (ensures correctness)

	sparrow.Eat()
	penguin.Eat() // ✅ Both birds can eat without issues
}
