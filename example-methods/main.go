package main

/*
 * Take a look to the example source https://medium.com/rungo/anatomy-of-methods-in-go-f552aaa8ac4a
 * Implementing interfaces
 */

import (
	"fmt"
	"my-go-examples/example-methods/nestedstruct"
	"my-go-examples/example-methods/pointerreceivers"
	"my-go-examples/example-methods/samename"
)

type employee struct {
	firstName, lastName string
}

func getFullNameFunction(firstName string, lastName string) (fullName string) {
	fmt.Println("Full name build with an simple function")
	fullName = firstName + " " + lastName
	return
}

func (e employee) getFullNameMethod() string {
	fmt.Println("Full name build with an employee method")
	return e.firstName + " " + e.lastName
}

func main() {
	ross := employee{
		firstName: "Ross",
		lastName:  "Geller",
	}

	fmt.Println()
	fmt.Println("A simple function")
	fmt.Println(getFullNameFunction(ross.firstName, ross.lastName))

	fmt.Println()
	fmt.Println("Switch to a method")
	fmt.Println(ross.getFullNameMethod())

	rectangle := samename.Rectangle{5.0, 4.0}
	circle := samename.Circle{5.0}
	fmt.Printf("Area of rectangle is %0.2f\n", rectangle.Area())
	fmt.Printf("Area of circle is %0.2f\n", circle.Area())

	fmt.Println()
	fmt.Println("Switch to a method")
	fmt.Println(ross.getFullNameMethod())

	fmt.Println()
	fmt.Println("Method pointer receiver")

	newEmployee := pointerreceivers.Employee{
		Name:   "Ross Geller",
		Salary: 50000,
	}

	newEmployeePointer := &newEmployee

	fmt.Println("new employee before name change =", newEmployeePointer.Name)

	newEmployeePointer.ChangeName("Monica Geller")

	fmt.Println("new employee after name change =", newEmployeePointer.Name)

	fmt.Println()
	fmt.Println("Calling methods with pointer receiver on values")
	fmt.Println("new employee before name change =", newEmployee.Name)

	newEmployee.ChangeNameWithGoShorcuts("Rachel Green")

	fmt.Println("new employee after name change =", newEmployee.Name)

	fmt.Println()
	fmt.Println("Methods on nested struct")

	employeeWithNestedStruct := nestedstruct.Employee{
		Name:    "Ross Geller",
		Salary:  1200,
		Contact: nestedstruct.Contact{"011 8080 8080", "New Delhi, India"},
	}

	fmt.Println("employee with nested struct before phone change =", employeeWithNestedStruct)

	employeeWithNestedStruct.ChangePhone("222 1010 1222")

	fmt.Println("employee with nested struct after phone change =", employeeWithNestedStruct)
}
