package main

import "fmt"

func main(){
	fmt.Println("Hello world")
}












// package main

// import (
// 	"fmt"
// )

// // User represents a system user.
// // Industry Tip: Use structs to group related data.
// type User struct {
// 	ID   int
// 	Name string
// }

// // updateName is a standalone function that modifies a User.
// // It takes a pointer (*User) to ensure changes affect the original object.
// func updateName(u *User) {
// 	// 1. Safety Check (Industry Best Practice)
// 	// Always verify a pointer is not nil before accessing fields to avoid crashes.
// 	if u == nil {
// 		return
// 	}

// 	// 2. Automatic Dereferencing
// 	// In Go, you don't need to write (*u).Name.
// 	// The compiler allows u.Name for convenience.
// 	u.Name = "Updated Name"
// }

// func main() {
// 	// 3. Initialization
// 	user := User{
// 		ID:   1,
// 		Name: "Original Name",
// 	}

// 	fmt.Println("Before:", user.Name)

// 	// 4. Passing the Address
// 	// We use the '&' operator to pass the memory address of 'user'
// 	// rather than a copy of the entire struct.
// 	updateName(&user)

// 	fmt.Println("After: ", user.Name) // Output: "Updated Name"
// }

