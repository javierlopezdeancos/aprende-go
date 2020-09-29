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

/***************************************/
/*       What is an interface?         */
/***************************************/

// We talked a lot about the object and behavior in the structs and methods lessons.
// We also saw how a structure (and other types) can implement methods.
// An interface is another piece of a puzzle that brings Go close to the Object-Oriented programming paradigm.
// An interface is a collection of method signatures that a Type can implement (using methods).
// Hence interface defines (not declares) the behavior of the object (of the type Type).

/***************************************/
/*         Declaring interface         */
/***************************************/

// Like struct, we need to create a derived type to simplify interface declaration using the keyword interface.

/*
type Shape interface {
    Area() float64
    Perimeter() float64
}
*/

// In the above example, we have defined the Shape interface which has two methods Area and Perimeter
// that accepts no arguments and return float64 value.
// Any type that implements these methods (with exact method signatures) will also implement Shape interface.

// Since the interface is a type just like a struct, we can create a variable of its type.
// In the above case, we can create a variable s of type interface Shape.

// An interface has two types. The static type of an interface is the interface itself,
// for example Shape in the above program. An interface does not have a static value,
// rather it points to a dynamic value.

// A variable of an interface type can hold a value of a type that implements the interface.

// The value of that type becomes the dynamic value of the interface and that type becomes the dynamic
// type of the interface.

// From the above example, we can see that zero value and type of the interface is nil.
// This is because, at this moment, we have declared the variable s of type Shape but did not assign any value.

// When we use Println function from fmt package with interface argument,
// it points to the dynamic value of the interface and %T syntax in Printf function refers to the dynamic type
// of interface.

func main() {

	/****************************************/
	/*       Implementing interface         */
	/****************************************/

	// Let’s declare Area and Perimeter methods with signatures provided by the Shape interface.
	// Also, let’s create Shape struct and make it implement Shape interface

	// In the above program, we’ve created the Shape interface and the struct type Rect.
	// Then we defined methods like Area and Perimeter which belongs to Rect type, therefore Rect implemented
	// those methods.

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

	/****************************************/
	/*          Empty interface             */
	/****************************************/

	// When an interface has zero methods, it is called an empty interface.
	// This is represented by interface{}. Since the empty interface has zero methods,
	// all types implement this interface implicitly.

	fmt.Println()
	fmt.Println("Empty interface")
	emptyinterface.Explain("hello world")
	emptyinterface.Explain(52)

	/****************************************/
	/*       Multiple interfaces            */
	/****************************************/

	// A type can implement multiple interfaces

	c := multiplesinterfaces.Cube{Side: 3}

	var sm multiplesinterfaces.Shape = c
	var o multiplesinterfaces.Object = c

	fmt.Println()
	fmt.Println("Multiples interfaces")
	fmt.Println("Area: ", sm.Area())
	fmt.Println("Volume: ", o.Volume())

	var sta multiplesinterfaces.Shape = multiplesinterfaces.Cube{Side: 10}
	c = sta.(multiplesinterfaces.Cube)

	/**************************************/
	/*          Type assertion            */
	/**************************************/

	// We can find out the underlying dynamic value of an interface
	// using the syntax i.(Type) where i is a variable of type interface and Type is a type that implements
	// the interface. Go will check if dynamic type of i is identical to the Type and return the dynamic
	// value is possible.

	fmt.Println()
	fmt.Println("Type assertion")
	fmt.Println("Area: ", c.Area())
	fmt.Println("Volume: ", c.Volume())

	/**************************************/
	/*          Type assertion            */
	/**************************************/

	// We have seen an empty interface and its use.
	// Let’s rethink the explain function we saw earlier.
	// As argument type of explain function is an empty interface, we can pass any argument to it.

	// But if the argument is a string, we want to the explain function to print the result in the uppercase.
	// How can we make that happen?

	// We can use ToUpper function from strings package but since it only accepts a string argument,
	// we need to make sure from inside the explain function that dynamic type of empty interface i is
	// string while doing so.

	// This can be done using Type switch. The syntax for type switch is similar to type
	// assertion and it is i.(type) where i is interface and type is a fixed keyword.
	// Using this we can get the dynamic type of the interface instead of the dynamic value.

	fmt.Println()
	fmt.Println("Type switch")
	fmt.Println("Area: ", c.Area())
	fmt.Println("Volume: ", c.Volume())

	/**************************************/
	/*      Embedding interfaces          */
	/**************************************/

	// In Go, an interface cannot implement other interfaces or extend them,
	// but we can create a new interface by merging two or more interfaces. Let’s rewrite our Shape-Cube program.

	// This is possible because, like anonymously nested struct, all methods of nested interfaces
	// get promoted to parent interfaces.

	fmt.Println()
	fmt.Println("Embedding interfaces")
	cmi := mergeinterfaces.Cube{20}

	var smi mergeinterfaces.Shape = cmi
	var omi mergeinterfaces.Object = cmi
	var mmi mergeinterfaces.Material = cmi

	fmt.Printf("Dynamic type '%T' and value '%v' of interface mmi of static type Material\n", mmi, mmi)
	fmt.Printf("Dynamic type '%T' and value '%v' of interface smi of static type Shape\n", smi, smi)
	fmt.Printf("Dynamic type '%T' and value '%v' of interface omi of static type Object\n", omi, omi)

	/**************************************/
	/*    Pointer vs Value receiver       */
	/**************************************/

	// Will interface be ok with method accepting pointer receiver too

	fmt.Println()
	fmt.Println("Pointer vs value receivers")

	r := pointerreceivers.Rectangle{Width: 20, Height: 10}
	var spr pointerreceivers.Shape = &r

	area := spr.Area()
	perimeter := spr.Perimeter()

	fmt.Println("Area of rectangle is", area)
	fmt.Println("Perimeter of rectangle is", perimeter)

	/**************************************/
	/*      Interface comparison          */
	/**************************************/

	// Two interfaces can be compared with == and != operators.
	// Two interfaces are always equal if the underlying dynamic values are nil, which means,
	//  two nil interfaces are always equal, hence == operation returns true.

	// If the dynamic types of the interface are not comparable, like for example, slice, map and function,
	// or the concrete value of an interface is a complex data structure like slice or array that contains
	// these uncomparable values, then == or != operations will result in a runtime panic.

	// If one interface is nil, then == operation will always return false.

	fmt.Println()
	fmt.Println("Interface comparison")

	var i, j interface{}
	fmt.Println(i == j) // true

	/**************************************/
	/*        Use of interfaces           */
	/**************************************/

	// We have learned interfaces and we saw they can take different forms. That’s the definition of polymorphism.

	// Interfaces are very useful in case of functions and methods where you need argument of dynamic types,
	// like Println function which accepts any type of values.

	// When multiple types implement the same interface, it becomes easy to work with them.
	// Hence whenever we can use interfaces, we should.
}
