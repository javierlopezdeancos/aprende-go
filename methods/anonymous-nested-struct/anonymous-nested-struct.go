package anonymous_nested_struct

/**
 * In the previous lesson, we also learned about anonymous fields and field promotions.
 * In a nutshell, if a field of a struct an anonymous struct, the nested struct fields
 * will be promoted to the parent.
 */

// Contact struct
type Contact struct {
	Phone, Address string
}

// Employee struct
type Employee struct {
	Name   string
	Salary int
	Contact
}

// ChangePhone method with Employee pointer receiver
func (c *Employee) ChangePhone(newPhone string) {
	c.Phone = newPhone
}
