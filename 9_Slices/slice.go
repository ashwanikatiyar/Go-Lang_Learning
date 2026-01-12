package main

import "fmt"

/*
	What is a slice?

	-A slice is a lightweight, flexible wrapper around an array.
	-It references an underlying array and can grow or shrink dynamically.

*/

func main() {

	//Declaring Syntax
	var s []int
	fmt.Println(s == nil) //true
	//Just Declaring Slice and not Initializing makes it 'nil' ,
	//Init like this or using make
	//s := []int{1, 2, 3}
	//s := make([]int, length, capacity)

	//ShortHand Declaration
	s2 := []int{}
	s2 = append(s2, 1, 2, 3)
	fmt.Println("Short Hand Slice2 : ", s2)

	//Initializing 2D Slice
	s3 := make([][]int, 2)
	fmt.Println("2D Slice using make : ", s3)




    // Slice a slice
    sub := s[1:4]
    fmt.Println("Sub-slice:", sub)



	// Assigning a slice copies the header, but both slices refer to the same underlying array:

	a := []int{1, 2, 3}
	b := a
	b[0] = 100

	fmt.Println(a) // [100, 2, 3]
	// Changes in b affect a.

}
