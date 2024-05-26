package nested_interface

// Salaried interface
type Salaried interface {
	GetSalary() int
}

// Salary struct
type Salary struct {
	Basic     int
	Insurance int
	Allowance int
}

// GetSalary method
func (s Salary) GetSalary() int {
	return s.Basic + s.Insurance + s.Allowance
}

// Employee struct
type Employee struct {
	FirstName, LastName string
	Salary              Salaried
}
