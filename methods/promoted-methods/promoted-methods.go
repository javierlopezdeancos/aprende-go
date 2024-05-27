package promoted_methods

/**
 * Like promoted fields, methods implemented by the anonymously nested struct are also promoted
 * to the parent struct. As we saw in the previous example, Contact field is anonymously nested.
 * Hence we could access phone field of the inner struct on the parent.
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
func (e *Contact) ChangePhone(newPhone string) {
	e.Phone = newPhone
}
