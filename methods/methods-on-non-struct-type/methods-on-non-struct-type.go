package methods_on_non_struct_type

/**
* So far we have seen methods belonging to struct type but from the definition of the methods,
* it is a function that can belong to any type. Hence a method can receive any type as long as
* the type definition and method definition is in the same package.
**/

import (
	"strings"
)

// MyString type an alias to string native type
type MyString string

// ToUpperCase upper case the string
func (s MyString) ToUpperCase() string {
	normalString := string(s)
	return strings.ToUpper(normalString)
}
