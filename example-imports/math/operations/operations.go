package operations

// Operators to apply an operation
type Operators struct {
	A int
	B int
}

//Sum do the 2 numbers sum
func (o Operators) Sum() int {
	return o.A + o.B
}

//Mul do the 2 numbers multiplication
func (o Operators) Mul() int {
	return o.A * o.B
}
