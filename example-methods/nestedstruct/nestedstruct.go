package nestedstruct

/*
* We learned a great deal about the nested structure in the previous lesson.
* As a struct field also can be a struct, we can define a method on parent struct and
* access nested struct to do anything we want.
 */

// Contact struct
type Contact struct {
	Phone, Address string
}

// Employee struct
type Employee struct {
	Name    string
	Salary  int
	Contact Contact
}

// ChangePhone method with Employee pointer receiver
func (e *Employee) ChangePhone(newPhone string) {
	e.Contact.Phone = newPhone
}

// ChangePhone method with Contact nested structure receiver
func (c *Contact) ChangePhone(newPhone string) {
	c.Phone = newPhone
}
