package main

import "fmt"

func main(){

	//Standard way
	// var name string = "Go Lang"

	var price float32 = 55.2
	fmt.Println("Price : ",price)

	//Go can infer the type 
	// var name = "Go Lang"

	//Assigning later
	var age int

	age = 22

	fmt.Println("Age : ",age)

	//Short Hand
	name := "Go Lang"
	fmt.Println(name)
}