package main

import (
	"fmt"
	"my-go-examples/example-imports/math"
	"my-go-examples/example-imports/math/operations"
)

func main() {

	var a math.Arithmetic
	a = operations.Operators{A: 2, B: 3}
	fmt.Println(a.Sum())
	fmt.Println(a.Mul())
}
