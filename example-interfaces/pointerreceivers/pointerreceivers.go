package pointerreceivers

// Shape interface with Area method
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle type with Width and Height prop
type Rectangle struct {
	Width  float64
	Height float64
}

// Area implementation to Rect
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter implementation to Rect
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width * r.Height)
}
