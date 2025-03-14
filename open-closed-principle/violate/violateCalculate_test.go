package openclosedprinciple

import (
	"fmt"
	"testing"
)

// Rectangle struct implements Shape
type Rectangle struct {
	Width, Height float64
}

// Circle struct implements Shape
type Circle struct {
	Radius float64
}

// Function to calculate the total area of shapes (violating OCP)
func CalculateTotalArea(rectangles []Rectangle, circles []Circle) float64 {
	totalArea := 0.0

	for _, r := range rectangles {
		totalArea += r.Width * r.Height
	}

	for _, c := range circles {
		totalArea += 3.14 * c.Radius * c.Radius
	}

	return totalArea
}

func TestCalulate(t *testing.T) {
	rectangles := []Rectangle{{Width: 10, Height: 5}}
	circles := []Circle{{Radius: 7}}

	totalArea := CalculateTotalArea(rectangles, circles)

	fmt.Println("Total Area:", totalArea)
}
