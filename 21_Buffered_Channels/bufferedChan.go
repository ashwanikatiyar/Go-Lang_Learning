package main

import "fmt"

func main() {
	// 1. Create a buffered channel of strings with a capacity of 2
	messages := make(chan string, 2)

	// 2. Send data to the channel
	// Because the channel is buffered, these sends do not block
	messages <- "Hello"
	messages <- "Buffered Channel"
	// messages <- "3rd Message"  //!Deadlock

	// 3. Receive data from the channel
	fmt.Println(<-messages) // Output: Hello
	fmt.Println(<-messages) // Output: Buffered Channel

}
