package main

import "fmt"

/*
*What Is a Channel?
A channel is a typed pipe used by goroutines to communicate.
Channels allow goroutines to safely share data.

*Mental model
Goroutine A ───▶  CHANNEL  ───▶ Goroutine B

*Go philosophy
Share memory by communicating, not communicate by sharing memory

*What Is a Channel — in Plain English
A channel is a safe mailbox.
	-One goroutine puts data in
	-Another goroutine takes data out
	-Go runtime ensures safety

*/


func main() {
	//Declaring
	messageChan := make(chan string)

	//Passing data into channel
	messageChan <- "Ping"

	//Getting data from channel
	data := <- messageChan

	fmt.Println(data)
}
