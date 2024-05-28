package main

import (
	simpleFunctionRules "aprende-go/functions/simple-function-rules"
	"fmt"
)

// A function, in general, is a small piece of code that is dedicated to a perform particular task based on
// some input values. We create a function so that we can perform such an operation whenever we want from wherever
// we want by providing some input values.

// These input values provide extra information to a function and they are totally optional. A function’s execution
// may or may not return any result.

// In Go, a function is defined using func keyword.

func main() {

	/***********************************/
	/*       A simple function        */
	/*********************************/

	fmt.Println()
	fmt.Println("A simple function")
	simpleFunctionRules.DoSomeThing()

	/******************************************/
	/*       Function name convention        */
	/****************************************/

	// Go recommends writing function names in simple word or camelCase. Even under_score function names are valid,
	// but they are not idiomatic in Go.

	/**************************************/
	/*        Function parameters        */
	/************************************/

	// As discussed earlier, a function may take input values for its execution.
	// These input values are provided in a function call, called arguments.
	//  One or multiple arguments can also be passed to a function.

	fmt.Println()
	fmt.Println("Function parameters")
	fmt.Println()
	fmt.Println("Print greeting message")
	simpleFunctionRules.PrintUserName("Javi")

	fmt.Println()
	fmt.Println("Print sum of to integers (2, 4)")
	simpleFunctionRules.PrintAdd(2, 4)

	/****************************************/
	/*            Return value             */
	/**************************************/

	fmt.Println()
	fmt.Println("Return value")
	fmt.Println()
	fmt.Println("Return sum of to integers (3, 5)")
	fmt.Println(simpleFunctionRules.Add(3, 5))

	/****************************************/
	/*        Multiple return values       */
	/**************************************/

	fmt.Println()
	fmt.Println("Multiple return values")
	fmt.Println()
	fmt.Println()
	fmt.Println("Return sum and mult of to integers (2, 3)")
	add, mult := simpleFunctionRules.AddAndMult(2, 3)
	fmt.Println("Sum", add)
	fmt.Println("Mult", mult)

	/****************************************/
	/*        Named  return values         */
	/**************************************/

	fmt.Println()
	fmt.Println("Named  return values")
	fmt.Println()
	fmt.Println("Return sum and mult of to integers (6, 3)")
	div, multiply := simpleFunctionRules.DivAndMult(6, 3)
	fmt.Println("Div", div)
	fmt.Println("Mult", multiply)

	/****************************************/
	/*        Recursive function           */
	/**************************************/

	fmt.Println()
	fmt.Println("Recursive function")
	fmt.Println()
	fmt.Println("Return factorial 4")
	fact := simpleFunctionRules.GetFactorial(4)
	fmt.Println("fact", fact)

	/****************************************/
	/*         Defer keyword              */
	/**************************************/

	// defer is a keyword in Go that makes a function executes at the end of the execution of parent
	// function or when parent function hits return statement.
	fmt.Println()
	fmt.Println("Defer keyword")
	fmt.Println()
	fmt.Println("main started")
	defer simpleFunctionRules.SayDone()
	fmt.Println("main finished")

	/****************************************/
	/*        Function as type             */
	/**************************************/

	fmt.Println()
	fmt.Println("Function as type")
	fmt.Println()
	fmt.Printf("Type of function add is			%T\n", simpleFunctionRules.Add)
	fmt.Printf("Type of function subtract is		%T\n", simpleFunctionRules.Subtract)

	addResult := simpleFunctionRules.Calc(5, 3, simpleFunctionRules.Add)
	subResult := simpleFunctionRules.Calc(5, 3, simpleFunctionRules.Subtract)

	fmt.Println()
	fmt.Println("5 + 3 =", addResult)
	fmt.Println("5 - 3 =", subResult)

	addResultPlus := simpleFunctionRules.CalcPlus(6, 10, simpleFunctionRules.Add)
	subResultPlus := simpleFunctionRules.CalcPlus(6, 10, simpleFunctionRules.Subtract)

	fmt.Println()
	fmt.Println("6 + 10 =", addResultPlus)
	fmt.Println("6 - 10 =", subResultPlus)

	/**************************************************/
	/*    Function as value (anonymous function)     */
	/************************************************/

	fmt.Println()
	fmt.Println("Function as value (anonymous function)")
	fmt.Println()
	simpleFunctionRules.SayYeah()

	/**********************************************************/
	/*    Immediately-invoked function expression (IIFE)     */
	/********************************************************/

	fmt.Println()
	fmt.Println("Immediately-invoked function expression (IIFE))")
	fmt.Println()

	sumIife := func(a int, b int) int {
		return a + b
	}(3, 5)

	fmt.Println("5+3 =", sumIife)
}
