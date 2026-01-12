package main

import "fmt"

func main() {

	//Range for Iterating over a Slice
	s1 := make([]int, 4)
	s1[0], s1[1], s1[2], s1[3] = 1, 2, 3, 4 //Adding values

	//Iterating
	for index, value := range s1 {
		fmt.Println("Index : ", index, "  Value : ", value)
	}

	//Range for iterating over a Map
	m1 := make(map[string]int)
	m1["Age"], m1["Height"] = 40, 180

	//Iterating Keys and Values
	for key, value := range m1 {
		fmt.Println("Key : ", key, "  Value : ", value)
	}
	//Iterating just Values
	for _, value := range m1 {
		fmt.Println("Value : ", value)
	}

	//Range for Iterating over strings
	mystr := "Ashwani"
	for index, char := range mystr {
		fmt.Println("Index : ", index, " UnicodeToChar : ", string(char))
	}
	/*
	FYI
	-char is a rune (Unicode character)
	-index is the byte position
	*/

}
