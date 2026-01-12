package main

import "fmt"

//Maps are similar to Obj in JS , Dict in Python , HashTable in Java

func main() {
	
	//Declaring a Map
	var m1 map[string]int
	fmt.Println("Declared a Map m1",m1)

	//Initializing Map using Make
	m2 := make(map[string]int)
	fmt.Println("Initialized a Map m2 using Make Func",m2)
	

	//Adding values 
	m2["Age"] = 24
	m2["Height"] = 180
	fmt.Println("Added values in M2 \n",m2)


	//Retrieving Elements
	value := m2["Age"]
	fmt.Println("Retrieving Value of Age from M2 : ", value)
	//If the key doesn't exist, the zero value (e.g., 0 for int, 
	// "" for string , false for bool) is returned.


	//Checking if a Key Exists
	v , ok := m2["Height"]
	if(!ok){
		fmt.Println("Key Doesnt Exists")
	}else {
		fmt.Println("Key found , here is the value : ",v)
	}


	// Length of a Map
	fmt.Println(len(m2))

	//Iterating over a Map
	for key, value := range m2{
		fmt.Println("Key: ",key , "Value",value)
	}

	//Delete a Key
	delete(m2 , "Age")
	fmt.Println("Deleted Age key from M2: ",m2)

	//Clearing the map outright
	clear(m2)
	fmt.Println("Clear Map m2: ",m2)
}