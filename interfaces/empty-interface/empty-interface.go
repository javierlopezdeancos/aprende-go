package empty_interface

import (
	"fmt"
	"strings"
)

// Explain function to detect de interface dynamic type value
func Explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("I stored string", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("I stored int", i)
	default:
		fmt.Println("I stored something else", i)
	}
}
