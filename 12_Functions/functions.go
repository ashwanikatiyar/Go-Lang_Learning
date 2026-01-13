package main

import (
	"fmt"
)

// Normal Add function take int type variable and returning int type
func add(a int, b int) int {
	return a + b
}

// Returning mulitple values
func getLanguages() (string, string, bool) {
	return "Go-Lang", "TypeScript", true
}

// Returning a slice
func getArray() []int {
	s := make([]int, 4)
	return s
}

// Naked Return - It is only possible when you have named return parameters in your function signature
func checkIfAdmin(role string) (Permission string, err error) {
	if role == "Admin" {
		Permission = "Granted"
		return
	} else if role != "Admin" {
		Permission = "Unauthorized"
		return
	} else {
		return
	}
}

// Variadic Functions - Func with variable num of Arguments
func sum(nums ...int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}

// Function as a Parameter
func applyOperation(a int, b int, op func(int, int) int) int {
	return op(a, b)
}
// Define some functions to pass as Parameter
// func "add" already made above , now making multiply
func multiply(x, y int) int {
    return x * y
}//This is called parameter shorthand or "type grouping" (Using one int instead of 2). 







func main() {
	fmt.Println(add(5, 5))

	//Kinda Destructuring
	var1, var2, bool1 := getLanguages()
	fmt.Println(var1, var2, bool1)

	//getting Array from Func
	fmt.Println(getArray())

	//Naked return or Named Return Values
	fmt.Println(checkIfAdmin("Admin")) //output = Granted <nil>
	fmt.Println(checkIfAdmin("User"))  // output = Unauthorized <nil>

	// Variadic Functions
	fmt.Println("Variadic function : ", sum(1, 2, 3))
	fmt.Println("Variadic function : ", sum(1, 2, 3, 5, 7, 8))

	//Function as a Parameter
	fmt.Println(applyOperation(2 ,3 ,add))
	fmt.Println(applyOperation(2 ,3 ,multiply))

}
