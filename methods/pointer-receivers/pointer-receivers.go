package pointer_receivers

/**
 * When a method belongs to a type, its receiver receives a copy of the object on which it was called.
 * To verify that, we can create a method that mutates a struct it receives.
 * Let’s create a method changeName that changes the name field of an Employee struct.
 */

// Employee struct
type Employee struct {
	Name   string
	Salary int
}

// ChangeName method to change the name field in an Employee
func (e *Employee) ChangeName(newName string) {
	(*e).Name = newName
}

// ChangeNameWithGoShorcuts method to change the name field in an Employee
func (e *Employee) ChangeNameWithGoShorcuts(newName string) {
	e.Name = newName
}
