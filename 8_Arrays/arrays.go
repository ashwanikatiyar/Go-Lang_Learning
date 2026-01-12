package main

import "fmt"

func main() {
	/*
		-Arrays are fixed-size sequences of elements of the same type.
		-The size of the array is part of its type, meaning [5]int and [10]int are     different types.
		-Initializing empty arrays will give Zeroes , "" , false for int , string and 
		bool respectively 
	*/



	//Syntax for Declaring var arrayName[size] type
	var arr1 [3]int // An array of 3 integers, initialized with zero values
	length := len(arr1)

	fmt.Println(length)

	//Or (Remember , curly braces)
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	//2D Array
	arr3 := [2][2]int{{1,2},{3,4}}
	fmt.Println("2D Array : ",arr3)

	//Assigning an array copies the entire array:

	a := [3]int{1, 2, 3}
	b := a
	b[0] = 100
	fmt.Println(a) // [1, 2, 3]
	fmt.Println(b) // [100, 2, 3]
	//Changes to b do not affect a.

	//In Go, slices are more flexible than arrays and are used extensively:

	slice1 := []int{1, 2, 3} // Slice
	fmt.Println("Slice = ",slice1)




	//Revise Array:

	// Declare and initialize array
    var numbers [5]int = [5]int{1, 2, 3, 4, 5}

    // Access elements
    fmt.Println("First element:", numbers[0])

    // Modify element
    numbers[2] = 99

    // Length of array
    fmt.Println("Array length:", len(numbers))

    // Loop through array
    for i, v := range numbers {
        fmt.Printf("Index %d: %d\n", i, v)
    }

	//In the above loop two variables i and v are
	// being init , first var takes the index and
	//second takes the value. THis is fixed mechanism 
	// of Go where first var takes the index and
	//second takes the value
}
