package main

import "fmt"

//First read the main function

func readOnly(c <-chan int) {
	fmt.Println("Taking readonly channel")
}
func writeOnly(c chan<- int) {
	fmt.Println("Taking write Only channel")
}

func main() {
	//Unidirectional Channel

	// readChannel := make(<-chan string)
	// writeChannel := make(chan<- string)

	/*
		Unidirectional channels are used as function parameters.
		You wont see Unidirectional channel very often , they are
		mainly used in function parameters

		Lets see how functions use them

	*/

	intChannel := make(chan int)
	readOnly(intChannel)
	writeOnly(intChannel)

	//Go to concept.txt for understanding reading and writing
	//!---------------------------------------------------------

	//Reading and Writing
	myIntChannel := make(chan int)

	go func() {
		myIntChannel <- 10
	}()
	//*One way of reading value ,
	// fmt.Println(<-myIntChannel)

	//*Better way (using one value and bool)
	value, ok := <-myIntChannel
	fmt.Println(`Value : `, value, `Boolean: `, ok)

	//!Doing the following will give deadlock
	//myIntChannel <- 10    (directly doing this without a go routine
	// will give deadlock error)
	// Because An unbuffered channel cannot "store" a value;
	//!-----------------------------

	/*Now go back to Concept.txt and reading about closing a
	Channel

	//!---------------------------------------------------------

	Lets first read the value , then close the channel ,then
	again try to read the value
	*/

	myStringChannel := make(chan string) //Creating channel

	go func() { //Writing a value in channel
		myStringChannel <- "Ashwani"
	}()

	strVal, ok := <-myStringChannel                       //Reading the value and Bool
	fmt.Println("String Value : ", strVal, "Bool : ", ok) //printing them

	close(myStringChannel)
	fmt.Println("Channel Closed")

	//Again try to read value
	strVal, ok = <-myStringChannel                        //Reading the value and Bool
	fmt.Println("String Value : ", strVal, "Bool : ", ok) //printing them

	/*
		! Output
		String Value :  Ashwani Bool :  true
		Channel Closed
		String Value :   Bool :  false
	*/

	//Now go to Concept.go and read about For range clause in Channel

	//!---------------------------------------------------------
	//NIL Channel
	fmt.Println("NIL Channel")


	var myNilChannel chan int
	//*Reading from NIL Channel
	
	// fmt.Println(<-myNilChannel) 
	//!Output : fatal error: all goroutines are asleep - deadlock!

	//*Closing a NIL Channel

	close(myNilChannel)
	//!Output : panic: close of nil channel

	//!---------------------------------------------------------

	//Buffered Channels
	// fmt.Println("Buffered Channel")

	// strChannel := make(chan string , 2)

}
