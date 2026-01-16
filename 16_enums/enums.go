package main

import "fmt"

/*
An enum is:
	A fixed set of named values
	Representing states / categories / modes
	Examples:
		-OrderStatus: Pending, Paid, Shipped
		-UserRole: Admin, User
		-PaymentStatus: Success, Failed

Basic Enum Pattern in Go
	Go uses:
		-type
		-const
		-iota
*/

// First eg
type OrderStatus int

const (
	Pending OrderStatus = iota
	Paid
	Shipped
	Cancelled
)

func checkStatus(status OrderStatus) string {
	switch status {
	case Pending:
		return "Pending"
	case Paid:
		return "Paid"
	case Shipped:
		return "Shipped"
	case Cancelled:
		return "Cancelled"
	default:
		return "Unknown"
	}
}

// 2nd Eg
type Role string

const (
	Admin  Role = "Admin"
	User        = "user"
	Hacker      = "hacker"
)

func checkRole(role Role) {
	fmt.Println("Role is : ", role)
}

func main() {
	//1St eg
	fmt.Println(checkStatus(Pending))

	//2nd eg
	checkRole(Hacker)


}
