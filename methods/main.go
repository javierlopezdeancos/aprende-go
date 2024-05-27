package main

/*
 * Take a look to the example source https://medium.com/rungo/anatomy-of-methods-in-go-f552aaa8ac4a
 * Anatomy of methods in Go
 */

import (
	acceptBothPointerAndValue "aprende-go/methods/accept-both-pointer-and-value"
	anonymousNestedStruct "aprende-go/methods/anonymous-nested-struct"
	methodsOnNonStructType "aprende-go/methods/methods-on-non-struct-type"
	nestedStruct "aprende-go/methods/nested-struct"
	pointerReceivers "aprende-go/methods/pointer-receivers"
	promotedMethods "aprende-go/methods/promoted-methods"
	sameName "aprende-go/methods/same-name"
	"fmt"
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

	/**********************************/
	/*      A simple function         */
	/**********************************/

	// A method is nothing but a function, but it belongs to a certain type.
	// A method is defined with slightly different syntax than a normal function.
	// It required an additional parameter known as a receiver which is a type to which the function belongs.
	// This way, a method (function) can access the properties of the receiver it belongs to (like fields of a struct).

	fmt.Println()
	fmt.Println("A simple function")
	fmt.Println(getFullNameFunction(ross.firstName, ross.lastName))

	/**********************************/
	/*      Switch to a method        */
	/**********************************/

	// A method can solve this problem easily.
	// To convert a function to the method, we just need an extra receiver parameter in the function definition.
	// The syntax for defining a method is as follows.

	// As a method belongs to a receiver type and it becomes available on that type as a property,
	// we can call that method using Type.methodName(...)syntax. In the above program,
	// we have used e.fullName() to get the full name of an employee since fullName method belongs to Employee.

	fmt.Println()
	fmt.Println("Switch to a method")
	fmt.Println(ross.getFullNameMethod())

	/*****************************************/
	/*      Methods with the same name       */
	/*****************************************/

	// One major difference between functions and methods is we can have multiple methods with same name while no two
	// functions with the same name can be defined in a package.
	// We are allowed to create methods with same name as long as their receivers are different.
	// Let’s create two struct types Circle and Rectangle and create two methods of the same name
	// Area which calculates the area of their receiver.

	rectangle := sameName.Rectangle{5.0, 4.0}
	circle := sameName.Circle{5.0}
	fmt.Printf("Area of rectangle is %0.2f\n", rectangle.Area())
	fmt.Printf("Area of circle is %0.2f\n", circle.Area())

	/**********************************/
	/*    Method pointer receiver     */
	/**********************************/

	// So far, we have seen methods belong to a type. But a method can also belong to the pointer of a type.

	// When a method belongs to a type, its receiver receives a copy of the object on which it was called. To verify
	// that, we can create a method that mutates a struct it receives.

	fmt.Println()
	fmt.Println("Method pointer receiver")

	newEmployee := pointerReceivers.Employee{
		Name:   "Ross Geller",
		Salary: 50000,
	}

	newEmployeePointer := &newEmployee

	fmt.Println("new employee before name change =", newEmployeePointer.Name)

	newEmployeePointer.ChangeName("Monica Geller")

	fmt.Println("new employee after name change =", newEmployeePointer.Name)

	/*********************************************************/
	/*    Calling methods with pointer receiver on values    */
	/*********************************************************/

	// If you are wondering, do I always need to create a pointer to work with methods with pointer receiver
	// then Go already figured that out.

	fmt.Println()
	fmt.Println("Calling methods with pointer receiver on values")

	fmt.Println("new employee before name change =", newEmployee.Name)

	newEmployee.ChangeNameWithGoShorcuts("Rachel Green")

	fmt.Println("new employee after name change =", newEmployee.Name)

	/**********************************/
	/*   Methods on nested struct     */
	/********************************+*/

	// If you are wondering, do I always need to create a pointer to work with methods with pointer receiver
	// then Go already figured that out.

	fmt.Println()
	fmt.Println("Methods on nested struct")

	cotactEmployee := nestedStruct.Contact{
		Phone:   "011 8080 8080",
		Address: "New Delhi, India",
	}

	employeeWithNestedStruct := nestedStruct.Employee{
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

	/*********************************/
	/*   Anonymously nested struct   */
	/*********************************/

	// If a field of a struct an anonymous struct,
	// the nested struct fields will be promoted to the parent.

	fmt.Println()
	fmt.Println("Anonymously nested struct")

	employeeWithNAnonymouslyNestedStruct := anonymousNestedStruct.Employee{
		Name:   "Ross Geller",
		Salary: 1200,
		Contact: anonymousNestedStruct.Contact{
			Phone:   "011 8080 8080",
			Address: "New Delhi, India",
		},
	}

	fmt.Println("employee before phone change =", employeeWithNAnonymouslyNestedStruct)

	employeeWithNAnonymouslyNestedStruct.ChangePhone("777 1313 1444")

	fmt.Println("employee after phone change =", employeeWithNAnonymouslyNestedStruct)

	/************************/
	/*   Promoted methods   */
	/************************/

	// Like promoted fields, methods implemented by the anonymously nested struct
	// are also promoted to the parent struct. As we saw in the previous example,
	// Contact field is anonymously nested.
	// Hence we could access phone field of the inner struct on the parent.

	fmt.Println()
	fmt.Println("Promoted methods")

	employeeWithPromotedMethod := promotedMethods.Employee{
		Name:   "Ross Geller",
		Salary: 1200,
		Contact: promotedMethods.Contact{
			Phone:   "011 8080 8080",
			Address: "New Delhi, India",
		},
	}

	fmt.Println("employee before phone change =", employeeWithPromotedMethod)

	employeeWithPromotedMethod.ChangePhone("999 9393 1999")

	fmt.Println("employee after phone change =", employeeWithPromotedMethod)

	/*************************************************/
	/*   Methods can accept both pointer and value   */
	/*************************************************/

	// When a normal function has a parameter definition,
	// it will only accept the argument of the type defined by the parameter.
	// If you passed a pointer to the function which expects a value, it will not work.
	// This is also true when function accepts pointer but you are passing a value instead.

	fmt.Println()
	fmt.Println("Methods can accept both pointer and value")

	employeeAcceptingBothPointerAndValue := acceptBothPointerAndValue.Employee{
		Name:   "Ross Geller",
		Salary: 1200,
	}

	fmt.Println("employee before change =", employeeAcceptingBothPointerAndValue)

	// We call ChangeName over the value and not the pointer but go fix that under the hood
	employeeAcceptingBothPointerAndValue.ChangeName("Monica Geller")

	// We call ShowSalary over the pointer and not the value but go fix that under the hood
	(&employeeAcceptingBothPointerAndValue).ShowSalary()

	fmt.Println("employee after change =", employeeAcceptingBothPointerAndValue)

	/**********************************/
	/*   Methods on non struct type   */
	/**********************************/

	// So far we have seen methods belonging to struct type but from the definition of the methods,
	// it is a function that can belong to any type.
	// Hence a method can receive any type as long as the type definition and method definition is in the same package.
	// So far, we defined struct and method in the same main package, hence it worked.

	fmt.Println()
	fmt.Println("Methods on non struct type")
	str := methodsOnNonStructType.MyString("Hello World")
	fmt.Println(str.ToUpperCase())
}
