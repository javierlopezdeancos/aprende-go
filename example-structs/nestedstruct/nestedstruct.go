package nestedstruct

// Salary employee struct
type Salary struct {
	Basic     int
	Insurance int
	Allowance int
}

// Employee struct
type Employee struct {
	FirstName, LastName string
	Salary
	Bool bool
}
