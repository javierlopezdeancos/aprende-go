package function_fields

// In the above program, we have defined struct type Employee which has two string fields and one
// function field. Just for simplicity, we have created a derived function type FullNameType.

// While creating the struct, we need to make sure the field FullName follows the function type syntax.
// In the above case, we assigned it with an anonymous function.

// FullNameType method to get the employee full name
type FullNameType func(string, string) string

// Employee structure
type Employee struct {
	FirstName, LastName string
	FullName            FullNameType
}
