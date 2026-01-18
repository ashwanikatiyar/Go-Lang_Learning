package main

import (
	"fmt"
	"time"
)

/*
1️ What Is a Goroutine?
A goroutine is a lightweight concurrent function managed by the Go runtime.

!!!!!!!!!!!!!!!!!!!
-Every process that runs in Concurrently in Golang is a Go routine
-Go routines are light wait threads
-Every program has atleast one go routine called the main function
-When the main function is terminated then all the go routines
will also get terminated , means all go routine work under main
!!!!!!!!!!!!!!!!!!!


2️ Your First Goroutine

	func sayHello() {
		fmt.Println("Hello")
	}

	func main() {
		go sayHello()
		time.Sleep(time.Second)
	}

Important
!If main() exits, all goroutines stop
*This is why we used Sleep (temporary hack , we will use Wait Groups)

3️Goroutines Run Concurrently (Not Parallel by Default)

		go task1()
		go task2()

	-Concurrent → interleaved execution
	-Parallel → multiple CPU cores
Go decides this automatically.





*/



func main() {
	
	for i:= 0 ; i < 10 ; i++{

		go func(i int){
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 2) //Just for learning , we use Waitgroups for making the main function wait.
}