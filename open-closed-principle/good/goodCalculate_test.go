package openclosedprinciple

import (
	"fmt"
	"testing"
)

// Shape interface (open for extension)
type Shape interface {
	Area() float64
}

// Rectangle struct implements Shape
type Rectangle struct {
	Width, Height float64
}

// Circle struct implements Shape
type Circle struct {
	Radius float64
}

type Triangle struct {
	Base, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Now our function is open for extension but closed for modification
func GoodCalculateTotalArea(shapes []Shape) float64 {
	totalArea := 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}

func TestGoodCalculate(t *testing.T) {
	shapes := []Shape{
		Triangle{Base: 5, Height: 2},
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 7},
	}

	totalArea := GoodCalculateTotalArea(shapes)
	fmt.Println("Total Area:", totalArea)
}
