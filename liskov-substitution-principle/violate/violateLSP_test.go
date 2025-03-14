package liskovsubstitutionprinciple

import (
	"fmt"
	"testing"
)

// Bird interface
type Bird interface {
	Fly()
}

// Sparrow can fly
type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow is flying!")
}

// Penguin cannot fly, violating LSP
type Penguin struct{}

func (p Penguin) Fly() {
	panic("Penguins can't fly!") // ❌ Unexpected behavior!
}

// Function to make a bird fly
func MakeBirdFly(b Bird) {
	b.Fly() // ❌ If a Penguin is passed, it will cause a runtime error
}

func TestBadLSP(t *testing.T) {
	sparrow := Sparrow{}
	// penguin := Penguin{}

	MakeBirdFly(sparrow) // ✅ Works fine
	// MakeBirdFly(penguin) // ❌ Violates LSP (crashes the program)
}
