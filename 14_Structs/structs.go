package main

import (
	"fmt"
	"time"
)

/*
A struct in Go is a user-defined composite data type that allows you to group together zero or more related values of different data types into a single, logical unit
*/

type Order struct {
	Id        int
	Status    string
	CreatedAt time.Time
}

// Lets try to make a function which modifies a struct
func (o Order) changeStatus() Order {
	o.Status = "Received"
	return o
}

//Now lets make a function which actually modifies struct using Pointers

// func (o *Order) changeStatusWithPointers()
func (o *Order) changeStatusWithPointers() {
	//No need to write *o.Status (dereferencing) bcoz struct does it for you
	o.Status = "Received"
}

/*Go doesnt have classes , structs are mostly used in place of
classes and similar to Constructors in Classes we can make
constructors for Structs as well.*/

func NewOrderConstructor(id int, status string) *Order {
	createdOrder := Order{
		Id:     id,
		Status: status,
	}
	return &createdOrder
}

/*//Little different syntax , direct return

func NewOrderConstructor2(id int , status string) *Order{
	return &Order{
		Id : id,
		Status: status,
	}
}
*/

/*Embedding
In OOPs languages we have Inheritance , in Go we have embedding
*/

type Address struct {
	City string
	Pincode int 
}
type User struct{
	Name string
	Age int
	// Address Address //Nesting - Name and type. For access : user.Address.City
	Address 	//Embedding -Just type. For access : user.City
}


func main() {
	myOrder := Order{
		Id:     1,
		Status: "Shipped",
	}
	myOrder.CreatedAt = time.Now()
	//If you dont set any field , default value is Zero Value (0 for int , "" for string , false for bool 0.0 for float )
	fmt.Println(myOrder)

	//Let change the status using changeStatus
	fmt.Println(myOrder.changeStatus()) //Changed to "Received"
	fmt.Println(myOrder)                //Not actually changed
	myOrder.changeStatusWithPointers()
	fmt.Println("Using Pointers : ", myOrder)

	//Lets try creating an order using Constructor
	myOrder2 := NewOrderConstructor(3, "Shipped")
	fmt.Println("My Order2 created using Constructor", myOrder2)

	//Anonymous Structs -Useful for quick grouping or tests.
	language := struct {
		Name           string
		IsProficient bool
	}{
		"Go-Lang",
		true,
	}
	fmt.Println("Anonymous Struct : ", language)


	//Embedding
	newUser := User{
		"Ashwani",
		24,
		Address{
			"Kanpur",
			208010,
		},
	}
	fmt.Println("Embedded Struct : ",newUser)

	//Lets change in the New User
	newUser.Address.City = "Mumbai"
	fmt.Println("After changing : ",newUser)
}
