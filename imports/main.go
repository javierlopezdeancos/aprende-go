package main

import (
	"aprende-go/imports/math"
	"aprende-go/imports/math/operations"
	"fmt"
)

func main() {

	var a math.Arithmetic
	a = operations.Operators{A: 2, B: 3}
	fmt.Println(a.Sum())
	fmt.Println(a.Mul())
}
