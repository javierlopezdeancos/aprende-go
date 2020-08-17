package main

/*
 * Take a look to the example source https://medium.com/rungo/interfaces-in-go-ab1601159b3a
 * Implementing interfaces
 */

import (
	"fmt"
	"my-go-examples/example-interfaces/emptyinterface"
	"my-go-examples/example-interfaces/mergeinterfaces"
	"my-go-examples/example-interfaces/multiplesinterfaces"
	"my-go-examples/example-interfaces/pointerreceivers"
)

type shape interface {
	perimeter() float64
	area() float64
}

type rect struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rect) perimeter() float64 {
	return r.width * r.height
}

func (r rect) area() float64 {
	return (r.width * r.height) / 2
}

func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func (c circle) area() float64 {
	return 3.14 * (c.radius * c.radius)
}

func main() {
	var s shape

	s = rect{width: 3, height: 10}

	fmt.Println()
	fmt.Println("Rectangle")
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perimeter())

	s = circle{radius: 3}

	fmt.Println()
	fmt.Println("Circle")
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perimeter())

	c := multiplesinterfaces.Cube{Side: 3}

	var sm multiplesinterfaces.Shape = c
	var o multiplesinterfaces.Object = c

	fmt.Println()
	fmt.Println("Multiples interfaces")
	fmt.Println("Area: ", sm.Area())
	fmt.Println("Volume: ", o.Volume())

	var sta multiplesinterfaces.Shape = multiplesinterfaces.Cube{Side: 10}
	c = sta.(multiplesinterfaces.Cube)

	fmt.Println()
	fmt.Println("Type assertion")
	fmt.Println("Area: ", c.Area())
	fmt.Println("Volume: ", c.Volume())

	fmt.Println()
	fmt.Println("Type switch")
	fmt.Println("Area: ", c.Area())
	fmt.Println("Volume: ", c.Volume())

	fmt.Println()
	fmt.Println("Empty interface")
	emptyinterface.Explain("hello world")
	emptyinterface.Explain(52)

	fmt.Println()
	fmt.Println("Embedding interfaces")
	cmi := mergeinterfaces.Cube{20}

	var smi mergeinterfaces.Shape = cmi
	var omi mergeinterfaces.Object = cmi
	var mmi mergeinterfaces.Material = cmi

	fmt.Printf("Dynamic type '%T' and value '%v' of interface mmi of static type Material\n", mmi, mmi)
	fmt.Printf("Dynamic type '%T' and value '%v' of interface smi of static type Shape\n", smi, smi)
	fmt.Printf("Dynamic type '%T' and value '%v' of interface omi of static type Object\n", omi, omi)

	fmt.Println()
	fmt.Println("Pointer receivers")

	r := pointerreceivers.Rectangle{Width: 20, Height: 10}
	var spr pointerreceivers.Shape = &r

	area := spr.Area()
	perimeter := spr.Perimeter()

	fmt.Println("Area of rectangle is", area)
	fmt.Println("Perimeter of rectangle is", perimeter)
}
