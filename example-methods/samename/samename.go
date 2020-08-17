package samename

/**
 * We are allowed to create methods with same name as long as their receivers are different.
 * Let’s create two struct * types Circle and Rectangle and create two methods of the same name
 * Area which calculates the area of their receiver.
 */

import (
	"math"
)

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Area method to rectangle receiver
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area method to circle receive
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
