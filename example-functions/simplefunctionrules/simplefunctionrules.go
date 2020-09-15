package simplefunctionrules

import "fmt"

// DoSomeThing example to a simple function
func DoSomeThing() {
	fmt.Println("Hello World! print in a simple function")
}

// PrintUserName print the user name string that pass in argument
func PrintUserName(userName string) {
	fmt.Println("Hello " + userName)
}

// As discussed earlier, a function may take input values for its execution.
// These input values are provided in a function call, called arguments.
// One or multiple arguments can also be passed to a function.

// PrintAdd sum and print two integers
func PrintAdd(a int, b int) {
	c := a + b
	fmt.Println(c)
}

// A function can also return a value that can be printed or assigned to another variable.
// You can use shorthand parameter notation in case multiple parameters in succession are of the same data type.
// In case a function returns a value, you must specify the data type of a return value just after
// the function parameter parentheses.

// Add return sum of two integers
func Add(a, b int) int {
	return a + b
}

// Unlike other programming languages, Go can return multiple values from the function.
// In this case, we must specify return types of the values (just like above) inside parentheses
// just after the function parameter parentheses.

// AddAndMult return sum and mult of two integers
func AddAndMult(a, b int) (int, int) {
	return a + b, a * b
}

// Named return values are a great way to explicitly mention return variables in the function definition itself.
// These variables will be created automatically and made available inside the function body.
// You can change the values of these variables inside a function.
// A return statement at the end of the function is necessary to return named values.
// Go will automatically return these variables when the function hits the return statement.

// DivAndMult return sum and mult of two integers
func DivAndMult(a, b int) (div, mul int) {
	div = a / b
	mul = a * b

	return
}

// A function is called recursive when it calls itself from inside the body.
// A simple syntax for the recursive function is
// func r() {
//   r()
// }

// GetFactorial return the number factorial
func GetFactorial(num int) int {
	// n! = n*(n-1)! where n>0
	if num > 1 {
		return num * GetFactorial(num-1)
	}

	return 1 // 1! == 1
}

// SayDone print I am done
func SayDone() {
	fmt.Println("I am done and I am deferred")
}

// Subtract return the two numbers substract
func Subtract(a int, b int) int {
	return a - b
}

// Calc do calc function that pass in third argument
func Calc(a int, b int, f func(int, int) int) int {
	r := f(a, b)
	return r
}

type calcFunc func(int, int) int

// CalcPlus do calc function that pass in third argument
func CalcPlus(a int, b int, f calcFunc) int {
	r := f(a, b) // calling add(a,b) or substract(a,b)
	return r
}

// A function in Go can also be a value.
// This means you can assign a function to a variable.

// SayYeah print yeah yeah yeah
var SayYeah = func() {
	fmt.Println("Yeah Yeah Yeah")
}
