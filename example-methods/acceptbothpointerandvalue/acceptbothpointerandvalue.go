package acceptbothpointerandvalue

/**
* When a normal function has a parameter definition, it will only accept the argument of the type
* defined by the parameter. If you passed a pointer to the function which expects a value,
* it will not work. This is also true when function accepts pointer but you are passing a value instead.
* But when it comes to methods, that’s not a strict rule. We can define a method with value or pointer receiver and
* call it on pointer or value. Go does the job of type conversion behind the scenes
**/

import "fmt"

// Employee struct
type Employee struct {
	Name   string
	Salary int
}

// ChangeName method with employee pointer receiver
func (e *Employee) ChangeName(newName string) {
	e.Name = newName
}

// ShowSalary method with employee receiver
func (e Employee) ShowSalary() {
	e.Salary = 1500
	fmt.Println("Salary of e =", e.Salary)
}
