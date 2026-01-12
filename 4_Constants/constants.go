/*Constants are fixed values that are assigned at compile time and cannot be changed during program execution. They are useful for defining values like mathematical constants, configuration values, or other fixed data.*/

package main

func main() {

	//Declaring Constants
	const pi float32 = 3.14
	const greeting = "Hello"

	//Multiple Constants
	const (
		StatusOk          = 200
		StatusNotFound    = 404
		StatusServerError = 500
	)

	// Types vs Untyped
	const (
		age int = 22
		greet = "Hello"		
	)	

	//Constant Expressions
	//Constants can be assigned the result of constant expressions, which are evaluated at compile time.
	const (
		a = 22
		b = 18
		c = a + b
	)

	// No need to necessarily use them , like var need to be used if declared
}
