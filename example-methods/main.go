package main

/*
 * Take a look to the example source https://medium.com/rungo/anatomy-of-methods-in-go-f552aaa8ac4a
 * Anatomy of methods in Go
 */

import (
	"fmt"
	"my-go-examples/example-methods/acceptbothpointerandvalue"
	"my-go-examples/example-methods/anonymouslynestedstruct"
	"my-go-examples/example-methods/methodsonnonstructtype"
	"my-go-examples/example-methods/nestedstruct"
	"my-go-examples/example-methods/pointerreceivers"
	"my-go-examples/example-methods/promotedmethods"
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

	/***********************************/
	/*      A simple function         */
	/*********************************/
	fmt.Println()
	fmt.Println("A simple function")
	fmt.Println(getFullNameFunction(ross.firstName, ross.lastName))

	/***********************************/
	/*      Switch to a method        */
	/*********************************/
	fmt.Println()
	fmt.Println("Switch to a method")
	fmt.Println(ross.getFullNameMethod())

	/*******************************************/
	/*      Methods with the same name       */
	/****************************************/
	rectangle := samename.Rectangle{5.0, 4.0}
	circle := samename.Circle{5.0}
	fmt.Printf("Area of rectangle is %0.2f\n", rectangle.Area())
	fmt.Printf("Area of circle is %0.2f\n", circle.Area())

	/***********************************/
	/*    Method pointer receiver     */
	/*********************************/
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

	/***********************************/
	/*   Methods on nested struct     */
	/*********************************/
	fmt.Println()
	fmt.Println("Methods on nested struct")

	cotactEmployee := nestedstruct.Contact{
		Phone:   "011 8080 8080",
		Address: "New Delhi, India",
	}

	employeeWithNestedStruct := nestedstruct.Employee{
		Name:    "Ross Geller",
		Salary:  1200,
		Contact: cotactEmployee,
	}

	fmt.Println("employee with nested struct before phone change =", employeeWithNestedStruct)

	employeeWithNestedStruct.ChangePhone("222 1010 1222")

	fmt.Println("employee with nested struct after phone change =", employeeWithNestedStruct)
	fmt.Println("employee with nested struct before phone change directly =", employeeWithNestedStruct)

	employeeWithNestedStruct.Contact.ChangePhone("333 6060 1333")

	fmt.Println("employee with nested struct after phone change directly =", employeeWithNestedStruct)

	/**********************************/
	/*   Anonymously nested struct   */
	/******************************* */
	fmt.Println()
	fmt.Println("Anonymously nested struct")

	employeeWithNAnonymouslyNestedStruct := anonymouslynestedstruct.Employee{
		Name:   "Ross Geller",
		Salary: 1200,
		Contact: anonymouslynestedstruct.Contact{
			Phone:   "011 8080 8080",
			Address: "New Delhi, India",
		},
	}

	fmt.Println("employee before phone change =", employeeWithNAnonymouslyNestedStruct)

	employeeWithNAnonymouslyNestedStruct.ChangePhone("777 1313 1444")

	fmt.Println("employee after phone change =", employeeWithNAnonymouslyNestedStruct)

	/*************************/
	/*   Promoted methods   */
	/***********************/
	fmt.Println()
	fmt.Println("Promoted methods")

	employeeWithPromotedMethod := promotedmethods.Employee{
		Name:   "Ross Geller",
		Salary: 1200,
		Contact: promotedmethods.Contact{
			Phone:   "011 8080 8080",
			Address: "New Delhi, India",
		},
	}

	fmt.Println("employee before phone change =", employeeWithPromotedMethod)

	employeeWithPromotedMethod.ChangePhone("999 9393 1999")

	fmt.Println("employee after phone change =", employeeWithPromotedMethod)

	/**************************************************/
	/*   Methods can accept both pointer and value   */
	/************************************************/
	fmt.Println()
	fmt.Println("Methods can accept both pointer and value")

	employeeAcceptingBothPointerAndValue := acceptbothpointerandvalue.Employee{
		Name:   "Ross Geller",
		Salary: 1200,
	}

	fmt.Println("employee before change =", employeeAcceptingBothPointerAndValue)

	// We call ChangeName over the value and not the pointer but go fix that under the hood
	employeeAcceptingBothPointerAndValue.ChangeName("Monica Geller")

	// We call ShowSalary over the pointer and not the value but go fix that under the hood
	(&employeeAcceptingBothPointerAndValue).ShowSalary()

	fmt.Println("employee after change =", employeeAcceptingBothPointerAndValue)

	/***********************************/
	/*   Methods on non struct type   */
	/*********************************/
	fmt.Println()
	fmt.Println("Methods on non struct type")
	str := methodsonnonstructtype.MyString("Hello World")
	fmt.Println(str.ToUpperCase())
}
