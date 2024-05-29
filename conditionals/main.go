package main

// Go provides if, if-else, if-else if-else variants of if/else statement we are familiar with.
// It is used to check a condition, and execute some code when the condition is true or false.

import (
	"fmt"
)

func main() {
	/***********************************/
	/*       The If condition          */
	/***********************************/

	// Simple use of if condition is demonstrated below. Unlike most of the programming languages,
	// Go does not allow to wrap the condition inside parenthesis ().

	fmt.Println()
	fmt.Println("The if condition")

	var condition = true

	if condition {
		fmt.Println("condition met")
	}

	fmt.Println("if terminated")

	/***********************************/
	/*     The if else condition       */
	/***********************************/

	// In the case of if-else condition, if the if condition is false, it will execute the else block.
	fmt.Println()
	fmt.Println("The if else condition")

	a := 2

	if a > 10 {
		fmt.Println("condition met")
	} else {
		fmt.Println("condition did not meet")
	}

	fmt.Println("program terminated")

	/******************************************/
	/*     The if else if else condition      */
	/******************************************/

	fmt.Println()
	fmt.Println("The if else if condition")

	// The else if condition is used when you want to try for multiple conditions.
	// First if condition will be evaluated, if the condition meets (evaluated to true),
	// code inside if block will be executed and no further conditions will be evaluated.

	// But if the if condition did not meet, the next conditions in succession of the else if blocks will be
	// evaluated until one condition meets. If none of the conditions meet, else block will be evaluated.

	fruit := "orange"

	if fruit == "mango" {
		fmt.Println("fruit is mango")
	} else if fruit == "orange" {
		fmt.Println("fruit is orange")
	} else if fruit == "banana" {
		fmt.Println("fruit is banana")
	} else {
		fmt.Println("I don't know which fruit this is")
	}

	/******************************************/
	/*          Initial statement             */
	/******************************************/

	// You can execute an initial statement before if condition evaluates.

	//   if statement; condition {
	//     ...
	//   }

	fmt.Println()
	fmt.Println("Initial condition")

	if fruit := "banana"; fruit == "mango" {
		fmt.Println("fruit is mango")
	} else if fruit == "orange" {
		fmt.Println("fruit is orange")
	} else if fruit == "banana" {
		fmt.Println("fruit is banana")
	} else {
		fmt.Println("I don't know which fruit this is")
	}

	// fruit variable is unavailable here
	fmt.Println(fruit)

	/********************************************************/
	/*          The switch conditional statement            */
	/********************************************************/

	//  switch input {
	//    case value_1:
	//      ...
	//    case value_1:
	// 			...

	fmt.Println()
	fmt.Println("The switch conditional statement")

	finger := 2

	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	}

	/*****************************************/
	/*          The default case             */
	/*****************************************/

	// Sometimes, we need else kind of condition where we need to execute some piece of code when none
	// of the cases match the input value. In that case, default block is used

	fmt.Println()
	fmt.Println("The default case")

	finger = 6

	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("No fingers matched")
	}

	// The default case does not have to be the last case. It can be anywhere in the switch block.

	/*******************************************/
	/*          Multiple case values           */
	/*******************************************/

	// You can present multiple (,) comma-separated case values to match the input switch value.

	fmt.Println()
	fmt.Println("Multiple case values")

	letter := "i"

	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Letter is a vovel.")
	default:
		fmt.Println("Letter is not a vovel.")
	}

	/*******************************************/
	/*          The initial statement          */
	/*******************************************/

	// There is another variant of switch statement where you can add initial statement
	// before the input value just like if conditional statement.

	// switch statement; input {
	//		...
	// }

	fmt.Println()
	fmt.Println("The initial statement")

	switch letter := "i"; letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Letter is a vovel.")
	default:
		fmt.Println("Letter is not a vovel.")
	}

	/****************************************/
	/*   Expressionless switch statement    */
	/****************************************/

	// Input value in Go’s switch statement is optional.
	// When there is no input value to switch statement,
	// case will be an expression that returns a boolean value.

	fmt.Println()
	fmt.Println("Expressionless switch statement")

	number := 20

	switch {
	case number <= 5:
		fmt.Println("number is less than or equal to 5")
	case number > 5:
		fmt.Println("number is greater than 5")
	case number > 10:
		fmt.Println("number is greater than 10")
	case number > 15:
		fmt.Println("number is greater than 15")
	}

	// Since number is greater than 20, number > 5 condition will meet and number
	// is greater than 5 will be printed to the console.

	/****************************************/
	/*      The fallthrough statement       */
	/****************************************/

	// As normally, when a case block executes, no other cases are tried. In the above program,
	// the number is greater than 5, 10 and 15 but since number > 5 and first case meets that condition and only
	// that case block is executed. To execute another case blocks, we use fallthrough statement.

	fmt.Println()
	fmt.Println("The fallthrough statement")

	switch number := 20; {
	case number <= 5:
		fmt.Println("number is less than or equal to 5")
		fallthrough
	case number > 5:
		fmt.Println("number is greater than 5")
		fallthrough
	case number > 10:
		fmt.Println("number is greater than 10")
		fallthrough
	case number > 15:
		fmt.Println("number is greater than 15")
	}

	// When a case condition is evaluated and returns true or matches the input value, that case block is executed.
	// If the compiler finds fallthrough statement, the next case block is executed without even evaluating the case.

	// The fallthrough should be the last statement in a case block,
	// else Go compiler will throw fallthrough statement out of place error.
}