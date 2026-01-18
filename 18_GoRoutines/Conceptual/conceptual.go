package main

import (
	"fmt"
	"time"
)

func sayHello() {
	time.Sleep(time.Second *2)
	fmt.Println("Hello World")
}

func sayHi() {
	fmt.Println("Hii")
}

func main() {
	//1st Run
	// fmt.Println("Learning Go Routines")
	// sayHello()
	// sayHi()

	/*
	Output:
		Learning Go Routines
		Hello World
		Hii

	Sequential right ? what even if sayHello takes two 
	seconds of time still go doesnt execute next func
	(sayHi()) until current func (sayHello()) is completed.

	Now comment 1st run and see the second run to understand the concept
	of using Go routines
	*/

	//2nd Run
	// fmt.Println("Learning Go Routines")
	// go sayHello()
	// sayHi()

	/*
	Output 
		Learning Go Routines
		Hii

	What just happened ? where is the response of sayHello() func?

	actually as soon as the main function exits all the function and
	goroutines gets terminated. 

	Now lets make the main function wait for a while before
	getting terminated/exited
	*/

	//3rd Run
	// fmt.Println("Learning Go Routines")
	// go sayHello()
	// sayHi()
	// time.Sleep(time.Second *3)

	/*
	Everything got printed because main function exited after every
	func got executed completely
	*/


	//4th Run (multiple go routine)
	fmt.Println("Learning Go Routines")
	go sayHello()
	go sayHi()
	time.Sleep(time.Second *3)

	//In 4th run , the demonstration shows the independency of multiple go routines
}
