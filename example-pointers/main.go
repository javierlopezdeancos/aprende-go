package main

import "fmt"

func changeValue(p *int) {
	*p = 2
}

func changeArrayValue(p *[3]int) {
	//*p == original array `a`
	// *p[0] != (*p)[0]
	(*p)[0] *= 2
	(*p)[1] *= 3
	(*p)[2] *= 4

	// using shorthand syntax
	// (*p)[0] == p[0]
	// p[0] *= 2
	// p[1] *= 3
	// p[2] *= 4
}

func main() {
	a := 0x00
	b := 0x0A
	c := 0xFF

	fmt.Printf("variable a of type %T with value %v in hex is %X\n", a, a, a)
	fmt.Printf("variable b of type %T with value %v in hex is %X\n", b, b, b)
	fmt.Printf("variable c of type %T with value %v in hex is %X\n", c, c, c)

	// print the memory addresses that store the variables values
	fmt.Println("&a =", &a)
	fmt.Println("&b =", &b)
	fmt.Println("&c =", &c)

	// Create a simple pointer
	var pa *int
	fmt.Printf("pointer pa of type %T with value %v\n", pa, pa)

	// Create a variable `int` type and a pointer pd that point to her.
	d := 1
	var pd *int
	pd = &d

	fmt.Printf("pointer pa of type %T with value %v\n", pd, pd)

	// Shorthand format notation to the last example
	e := 1
	pe := &e

	fmt.Printf("pointer pa of type %T with value %v\n", pe, pe)

	// Get the pointer value that is pointing
	fmt.Printf("data at %v is %v\n", pe, *pe)

	// Changing the variable value using a Pointer
	f := 1
	pf := &f
	*pf = 2
	fmt.Printf("a = %v\n", f)
	fmt.Printf("data at %v is %v\n", pf, *pf)

	// The new function
	pg := new(int)
	fmt.Printf("data at %v is %v\n", pg, *pg)

	// Passing a Pointer to a Function
	g := 1
	changeValue(&g)

	fmt.Printf("g = %v\n", g)

	h := [3]int{1, 2, 3}
	changeArrayValue(&h)

	fmt.Printf("h = %v\n", h)
}
