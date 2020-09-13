package main

import (
	"fmt"
	"my-go-examples/example-structs/exportedfields"
	"my-go-examples/example-structs/functionfields"
	"my-go-examples/example-structs/nestedinterface"
	"my-go-examples/example-structs/nestedstruct"
	"my-go-examples/example-structs/promotednestedinterface"
)

/*
 * Take a look to the example source https://medium.com/rungo/structures-in-go-76377cc106a2
 * Structs in Go (structs)
 */

/***********************************/
/*       A simple struct          */
/*********************************/
type employee struct {
	firstName string
	lastName  string
	salary    int
	fullTime  bool
}

func main() {

	/***********************************/
	/*       Create a struct          */
	/*********************************/
	var ross employee

	fmt.Println()
	fmt.Println("Create a struct")
	fmt.Println(ross)

	/*************************************************/
	/*      Getting and setting struct fields       */
	/***********************************************/
	fmt.Println()
	fmt.Println("Getting and setting struct fields")

	var phoebe employee
	phoebe.firstName = "Phoebe"
	phoebe.lastName = "Buffet"
	phoebe.salary = 1200
	phoebe.fullTime = true

	fmt.Println("phoebe.firstName =", ross.firstName)
	fmt.Println("phoebe.lastName =", ross.lastName)
	fmt.Println("phoebe.salary =", ross.salary)
	fmt.Println("phoebe.fullTime =", ross.fullTime)

	/************************************/
	/*      Initializing a struct      */
	/**********************************/
	fmt.Println()
	fmt.Println("Initializing a struct")

	rachel := employee{
		firstName: "Rachel",
		lastName:  "Green",
		fullTime:  true,
		salary:    600,
	}

	// There is one other way of initializing a struct that does not include field name declarations
	// rachel := employee{"Rachel", "Green", true, 600}

	fmt.Println(rachel)

	/************************************/
	/*       Anonymous struct          */
	/**********************************/
	fmt.Println()
	fmt.Println("Anonymous struct")

	monica := struct {
		firstName, lastName string
		salary              int
		fullTime            bool
	}{
		firstName: "Monica",
		lastName:  "Geller",
		salary:    1200,
	}

	fmt.Println(monica)

	/**************************************/
	/*       Pointer to  a struct        */
	/************************************/
	fmt.Println()
	fmt.Println("Pointer to a struct")

	lukePointer := &employee{
		firstName: "Luke",
		lastName:  "Skywalker",
		salary:    2200,
		fullTime:  true,
	}

	// since lukePointer is a pointer, we need to use *lukePointer dereferencing syntax to get the actual value
	// of the struct it is pointing to and the use
	fmt.Println("firstName: ", (*lukePointer).firstName)

	// But Go provide an easy alternative syntax to access fields. We can access the fields of a struct
	// pointer without dereferencing it first. Go will take care of dereferencing a pointer under the hood.
	fmt.Println("lastName: ", lukePointer.lastName)

	/**************************************/
	/*        Anonymous fields           */
	/************************************/
	fmt.Println()
	fmt.Println("Anonymous fields")

	// You can define a struct type without declaring any field names.
	// You have to just define the field data types and Go will use the data type declarations (keywords)
	// as the field names.

	type jedi struct {
		string
		int
		bool
	}

	lukeSkywalker := jedi{"Luke Skywalker", 32, true}
	lukeSkywalker.bool = false
	fmt.Println(lukeSkywalker.string, lukeSkywalker.int, lukeSkywalker.bool)

	/**************************************/
	/*          Nested struct            */
	/************************************/
	fmt.Println()
	fmt.Println("Nested struct")

	// A struct field can be of any data type. Hence, it is perfectly legal to have a struct
	// field that holds another struct. Hence, a struct field can have a data type that is a struct type.
	//  When a struct field has a struct value, that struct value is called a nested struct since it is nested
	// inside a parent struct.

	fernando := nestedstruct.Employee{
		FirstName: "Fernando",
		LastName:  "Redondo",
		Bool:      true,
		Salary:    nestedstruct.Salary{1100, 50, 50},
	}

	fmt.Println(fernando)

	/**************************************/
	/*        Promoted fields            */
	/************************************/
	fmt.Println()
	fmt.Println("Promoted fields")

	// We have learned that it is perfectly legal to define a struct type without declaring the field names
	// and Go will define the field names from the field types. This approach can also be applied in the nested struct.
	// We can drop the field name of a nested struct and Go will use struct type as the field name.

	fmt.Println("Fernando Redondo's basic salary", fernando.Salary.Basic)

	// When we use an anonymous nested struct, all the nested struct fields are automatically available on parent
	// struct. This is called field promotion.
	// Only the non-conflicting fields will get promoted.

	fmt.Println("Fernando Redondo's insurance salary", fernando.Insurance)

	/**************************************/
	/*        Nested interface           */
	/************************************/
	fmt.Println()
	fmt.Println("Nested interface")

	// Since we know that, an interface type is a declaration of method signatures.
	// Any data type that implements an interface can also be represented as a type of that interface (polymorphism).
	// We can have a struct field of an interface type and its value can be anything that implements that interface.

	gento := nestedinterface.Employee{
		FirstName: "Paco",
		LastName:  "Gento",
		Salary:    nestedinterface.Salary{1100, 50, 50},
	}

	fmt.Println("Paco Gento's salary is", gento.Salary.GetSalary())

	/******************************************/
	/*        Anonymous nested interface     */
	/****************************************/
	fmt.Println()
	fmt.Println("Anonymous nested interface")

	// Similar to the field promotions we saw earlier, methods are also promoted when a struct field
	// is an anonymous interface.

	maradona := promotednestedinterface.Employee{
		FirstName: "Diego",
		LastName:  "Armando",
		Salaried:  promotednestedinterface.Salary{12349100, 23450, 50},
	}

	fmt.Println("Diego Armando's salary is", maradona.GetSalary())

	/*******************************/
	/*      Exported fields       */
	/*****************************/
	fmt.Println()
	fmt.Println("Exported fields")

	// We can't initialize salary and fulltime properties beacuse they aren't exported
	rafael := exportedfields.Employee{
		FirstName: "Rafael",
		LastName:  "Martin Vazquez",
	}

	fmt.Println(rafael)

	/*******************************/
	/*      Function fields       */
	/*****************************/
	fmt.Println()
	fmt.Println("Function fields")

	// Struct fields can also be functions.

	sidi := functionfields.Employee{
		FirstName: "Ruiz",
		LastName:  "Diaz de Vivar",
		FullName: func(firstName string, lastName string) string {
			return firstName + " " + lastName
		},
	}

	fmt.Println(sidi.FullName(sidi.FirstName, sidi.LastName))

	/************************************/
	/*      Function comparasion       */
	/**********************************/
	fmt.Println()
	fmt.Println("Function comparasion")

	// Two structs are comparable if they belong to the same type and have the same field values.

	hugo := employee{
		firstName: "Hugo",
		lastName:  "Sanchez",
		salary:    1200,
	}

	hugoSanchez := employee{
		firstName: "Hugo",
		lastName:  "Sanchez",
		salary:    1200,
	}

	fmt.Println(hugo == hugoSanchez)

	/**************************************/
	/*      Struct field meta-data       */
	/************************************/
	fmt.Println()
	fmt.Println("Struct field meta-data")

	// Struct gives one more ability to add meta-data to its fields. Usually,
	// it is used to provide transformation information on how a struct field is encoded to or decoded
	// from another format (or stored/retrieved from a database), but you can use it to store whatever meta-info you
	// want to, either intended for another package or for your own use.

	// This meta-information is defined by the string literal (read strings examples).

	/*
		type Employee struct {
			firstName string `json:"firstName"`
			lastName  string `json:"lastName"`
			salary    int    `json: "salary"`
			fullTime  int    `json: "fullTime"`
		}
	*/
}
