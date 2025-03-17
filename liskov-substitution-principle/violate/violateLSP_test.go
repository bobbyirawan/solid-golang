package violate

import (
	"fmt"
	"testing"
)

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Method to set width
func (r *Rectangle) SetWidth(w float64) {
	r.Width = w
}

// Method to set height
func (r *Rectangle) SetHeight(h float64) {
	r.Height = h
}

// Method to calculate area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Square struct inherits Rectangle
type Square struct {
	Rectangle
}

// Override SetWidth and SetHeight to maintain square properties
func (s *Square) SetWidth(w float64) {
	s.Width = w
	s.Height = w // Force height to be same as width
}

func (s *Square) SetHeight(h float64) {
	s.Width = h
	s.Height = h // Force width to be same as height
}

func TestLSP(t *testing.T) {
	//var r Rectangle = Square{} // ❌ LSP Violation!
	// r.SetWidth(10)
	// r.SetHeight(20)

	fmt.Println("Expected Area: 10 * 20 = 200")
	// fmt.Println("Actual Area:", r.Area()) // ❌ Incorrect output!
}
