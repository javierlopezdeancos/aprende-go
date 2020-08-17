package mergeinterfaces

// Shape interface with Area method
type Shape interface {
	Area() float64
}

// Object interface with Volume method
type Object interface {
	Volume() float64
}

// Material interface merge between Object and Shape interfaces
type Material interface {
	Shape
	Object
}

// Cube type with Side prop
type Cube struct {
	Side float64
}

// Area implementation to Cube
func (c Cube) Area() float64 {
	return 6 * (c.Side * c.Side)
}

// Volume implementation to Cube
func (c Cube) Volume() float64 {
	return c.Side * c.Side * c.Side
}
